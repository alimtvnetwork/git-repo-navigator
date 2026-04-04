package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/user/gitmap/constants"
	"github.com/user/gitmap/model"
)

// runTaskRun starts the file-sync watch loop for a named task.
func runTaskRun(args []string) {
	fs := flag.NewFlagSet("task-run", flag.ExitOnError)

	var interval int
	var verbose, dryRun bool

	fs.IntVar(&interval, constants.FlagTaskInterval, constants.TaskDefaultInterval, constants.FlagDescTaskInterval)
	fs.BoolVar(&verbose, constants.FlagTaskVerbose, false, constants.FlagDescTaskVerbose)
	fs.BoolVar(&dryRun, constants.FlagTaskDryRun, false, constants.FlagDescTaskDryRun)
	fs.Parse(args)

	name := fs.Arg(0)
	if name == "" {
		fmt.Fprint(os.Stderr, constants.ErrTaskNameRequired)
		os.Exit(1)
	}

	interval = enforceMinInterval(interval)
	tasks := loadTaskFile()
	entry := findTaskByName(tasks, name)

	fmt.Printf(constants.MsgTaskRunning, name, interval)
	runSyncLoop(entry, interval, verbose, dryRun)
}

// enforceMinInterval clamps the interval to the minimum.
func enforceMinInterval(interval int) int {
	if interval < constants.TaskMinInterval {
		return constants.TaskMinInterval
	}

	return interval
}

// runSyncLoop runs the sync cycle on a timer until interrupted.
func runSyncLoop(entry model.TaskEntry, interval int, verbose, dryRun bool) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	syncOnce(entry, verbose, dryRun)

	for {
		select {
		case <-sigChan:
			fmt.Printf(constants.MsgTaskStopped, entry.Name)

			return
		case <-ticker.C:
			syncOnce(entry, verbose, dryRun)
		}
	}
}

// syncOnce performs a single sync pass from source to destination.
func syncOnce(entry model.TaskEntry, verbose, dryRun bool) {
	ignorePatterns := loadGitignorePatterns(entry.Source)
	syncCount := 0

	err := filepath.Walk(entry.Source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		relPath := relativePath(entry.Source, path)
		if isIgnored(relPath, info.IsDir(), ignorePatterns) {
			if info.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		if info.IsDir() {
			return nil
		}

		synced := syncSingleFile(entry.Dest, relPath, info, dryRun, verbose)
		if synced {
			syncCount++
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, constants.ErrTaskSyncFailed, entry.Source, err)
	}

	if syncCount == 0 && verbose {
		fmt.Print(constants.MsgTaskUpToDate)
	}
}

// syncSingleFile compares timestamps and copies if source is newer.
func syncSingleFile(destRoot, relPath string, srcInfo os.FileInfo, dryRun, verbose bool) bool {
	destPath := filepath.Join(destRoot, relPath)
	destInfo, err := os.Stat(destPath)

	isNewer := err != nil || srcInfo.ModTime().After(destInfo.ModTime())
	if isNewer {
		return handleSyncCopy(destRoot, relPath, srcInfo, dryRun, verbose)
	}

	return false
}

// handleSyncCopy performs the copy or dry-run print.
func handleSyncCopy(destRoot, relPath string, srcInfo os.FileInfo, dryRun, verbose bool) bool {
	if dryRun {
		fmt.Printf(constants.MsgTaskDrySync, relPath)

		return true
	}

	srcPath := filepath.Join(filepath.Dir(destRoot), relPath)
	destPath := filepath.Join(destRoot, relPath)

	err := ensureDestDir(destPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, constants.ErrTaskDestCreate, err)

		return false
	}

	err = copyFileContent(srcPath, destPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, constants.ErrTaskSyncFailed, relPath, err)

		return false
	}

	if verbose {
		fmt.Printf(constants.MsgTaskSynced, relPath)
	}

	return true
}

// ensureDestDir creates parent directories for a destination file.
func ensureDestDir(destPath string) error {
	return os.MkdirAll(filepath.Dir(destPath), constants.DirPermission)
}

// copyFileContent copies file content from source to destination.
func copyFileContent(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	buf := make([]byte, constants.TaskCopyBufferSize)
	_, err = io.CopyBuffer(destFile, srcFile, buf)

	return err
}

// relativePath returns a path relative to the base directory.
func relativePath(base, path string) string {
	rel, err := filepath.Rel(base, path)
	if err != nil {
		return path
	}

	return rel
}

// loadGitignorePatterns reads .gitignore patterns from a directory.
func loadGitignorePatterns(dir string) []string {
	path := filepath.Join(dir, ".gitignore")
	data, err := os.ReadFile(path)

	if err != nil {
		return nil
	}

	return parseGitignoreLines(string(data))
}

// parseGitignoreLines extracts patterns from gitignore content.
func parseGitignoreLines(content string) []string {
	var patterns []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if isGitignoreComment(trimmed) {
			continue
		}

		patterns = append(patterns, trimmed)
	}

	return patterns
}

// isGitignoreComment returns true for empty lines and comments.
func isGitignoreComment(line string) bool {
	if line == "" {
		return true
	}

	return strings.HasPrefix(line, "#")
}

// isIgnored checks if a path matches any gitignore pattern.
func isIgnored(relPath string, isDir bool, patterns []string) bool {
	if len(patterns) == 0 {
		return false
	}

	var wg sync.WaitGroup
	result := make(chan bool, len(patterns))

	for _, pattern := range patterns {
		wg.Add(1)

		go func(p string) {
			defer wg.Done()

			if matchesPattern(relPath, isDir, p) {
				result <- true
			}
		}(pattern)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for matched := range result {
		if matched {
			return true
		}
	}

	return false
}

// matchesPattern checks if a path matches a single gitignore pattern.
func matchesPattern(relPath string, isDir bool, pattern string) bool {
	isDirPattern := strings.HasSuffix(pattern, "/")
	if isDirPattern && isDir {
		cleanPattern := strings.TrimSuffix(pattern, "/")

		return matchGlob(relPath, cleanPattern)
	}

	if isDirPattern {
		return false
	}

	return matchGlob(relPath, pattern)
}

// matchGlob performs glob matching against path components.
func matchGlob(relPath, pattern string) bool {
	matched, err := filepath.Match(pattern, filepath.Base(relPath))
	if err != nil {
		return false
	}

	return matched
}
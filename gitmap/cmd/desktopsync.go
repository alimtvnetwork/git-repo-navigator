package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/user/gitmap/constants"
	"github.com/user/gitmap/model"
)

// runDesktopSync handles the "desktop-sync" subcommand.
func runDesktopSync() {
	outputDir := constants.DefaultOutputFolder
	if dirMissing(outputDir) {
		fmt.Fprintln(os.Stderr, constants.MsgNoOutputDir)
		os.Exit(1)
		return
	}
	jsonPath := filepath.Join(outputDir, constants.DefaultJSONFile)
	if fileMissing(jsonPath) {
		fmt.Fprintf(os.Stderr, constants.MsgNoJSONFile, jsonPath)
		os.Exit(1)
		return
	}
	records := loadRecords(jsonPath)
	syncToDesktop(records, jsonPath)
}

// dirMissing returns true if the directory does not exist.
func dirMissing(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return true
	}

	return !info.IsDir()
}

// fileMissing returns true if the file does not exist.
func fileMissing(path string) bool {
	_, err := os.Stat(path)

	return err != nil
}

// loadRecords reads and parses the JSON file into ScanRecords.
func loadRecords(path string) []model.ScanRecord {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", path, err)
		os.Exit(1)
	}
	var records []model.ScanRecord
	err = json.Unmarshal(data, &records)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON from %s: %v\n", path, err)
		os.Exit(1)
	}

	return records
}

// syncToDesktop registers each repo with GitHub Desktop.
func syncToDesktop(records []model.ScanRecord, source string) {
	if desktopMissing() {
		fmt.Fprintln(os.Stderr, constants.MsgDesktopNotFound)
		os.Exit(1)
		return
	}
	fmt.Printf(constants.MsgDesktopSyncStart, source)
	added, skipped, failed := syncAll(records)
	fmt.Printf(constants.MsgDesktopSyncDone, added, skipped, failed)
}

// desktopMissing returns true if GitHub Desktop CLI is not found.
func desktopMissing() bool {
	_, err := exec.LookPath(constants.GitHubDesktopBin)

	return err != nil
}

// syncAll iterates records and syncs each to GitHub Desktop.
func syncAll(records []model.ScanRecord) (added, skipped, failed int) {
	for _, r := range records {
		result := syncOne(r)
		added, skipped, failed = tallyResult(result, added, skipped, failed)
	}

	return added, skipped, failed
}

// syncResult represents the outcome of syncing one repo.
type syncResult int

const (
	syncAdded   syncResult = iota
	syncSkipped
	syncFailed
)

// syncOne attempts to register a single repo with GitHub Desktop.
func syncOne(r model.ScanRecord) syncResult {
	repoPath := r.AbsolutePath
	if repoPath == "" {
		fmt.Printf(constants.MsgDesktopSyncFailed, r.RepoName, "no absolute path")

		return syncFailed
	}
	if pathMissing(repoPath) {
		fmt.Printf(constants.MsgDesktopSyncSkipped, r.RepoName)

		return syncSkipped
	}

	return registerOne(r.RepoName, repoPath)
}

// pathMissing returns true if a directory path does not exist.
func pathMissing(path string) bool {
	_, err := os.Stat(path)

	return err != nil
}

// registerOne calls the GitHub Desktop CLI for a single repo.
func registerOne(name, repoPath string) syncResult {
	cmd := exec.Command(constants.GitHubDesktopBin, repoPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf(constants.MsgDesktopSyncFailed, name, fmt.Sprintf("%v: %s", err, output))

		return syncFailed
	}
	fmt.Printf(constants.MsgDesktopSyncAdded, name)

	return syncAdded
}

// tallyResult increments the appropriate counter.
func tallyResult(r syncResult, added, skipped, failed int) (int, int, int) {
	if r == syncAdded {
		return added + 1, skipped, failed
	}
	if r == syncSkipped {
		return added, skipped + 1, failed
	}

	return added, skipped, failed + 1
}

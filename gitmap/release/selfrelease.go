package release

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/user/gitmap/constants"
	"github.com/user/gitmap/store"
)

// ExecuteSelf resolves the gitmap source repo from the running binary
// (with DB fallback), switches to that directory, runs Execute,
// then returns to the original dir.
func ExecuteSelf(opts Options) error {
	srcRoot, err := resolveSourceRepo()
	if err != nil {
		return err
	}

	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not determine current directory: %w", err)
	}

	// Skip directory switch if already in the source repo.
	cleanSrc := filepath.Clean(srcRoot)
	cleanCwd := filepath.Clean(originalDir)

	if cleanSrc == cleanCwd {
		fmt.Printf(constants.MsgSelfReleaseSameDir, srcRoot)

		return Execute(opts)
	}

	fmt.Printf(constants.MsgSelfReleaseSwitch, srcRoot)

	err = os.Chdir(srcRoot)
	if err != nil {
		return fmt.Errorf("could not switch to source repo: %w", err)
	}

	releaseErr := Execute(opts)

	// Always attempt to return to original directory.
	if cdErr := os.Chdir(originalDir); cdErr != nil {
		fmt.Fprintf(os.Stderr, "  ⚠ Could not return to %s: %v\n", originalDir, cdErr)
	} else {
		fmt.Printf(constants.MsgSelfReleaseReturn, originalDir)
	}

	return releaseErr
}

// IsInsideGitRepo checks if the current directory is inside a Git repository.
func IsInsideGitRepo() bool {
	dir, err := os.Getwd()
	if err != nil {
		return false
	}

	return findGitRoot(dir) != ""
}

// resolveSourceRepo finds the git root of the gitmap source.
// It tries the executable path first, then falls back to the DB setting.
func resolveSourceRepo() (string, error) {
	// Strategy 1: resolve from executable path.
	if root, err := resolveFromExecutable(); err == nil && root != "" {
		saveSourceRepoDB(root)

		return root, nil
	}

	// Strategy 2: load from database.
	if root := loadSourceRepoDB(); root != "" {
		// Verify the path still has a .git directory.
		if findGitRoot(root) != "" {
			return root, nil
		}
	}

	return "", fmt.Errorf("%s", constants.ErrSelfReleaseNoRepo)
}

// resolveFromExecutable walks up from the binary location to find a .git root.
func resolveFromExecutable() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf(constants.ErrSelfReleaseExec, err)
	}

	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		return "", fmt.Errorf(constants.ErrSelfReleaseExec, err)
	}

	root := findGitRoot(filepath.Dir(exe))
	if root == "" {
		return "", fmt.Errorf("%s", constants.ErrSelfReleaseNoRepo)
	}

	return root, nil
}

// saveSourceRepoDB persists the source repo path in the Settings table.
func saveSourceRepoDB(path string) {
	db, err := store.OpenDefault()
	if err != nil {
		return
	}
	defer db.Close()

	_ = db.SetSetting(constants.SettingSourceRepoPath, path)
}

// loadSourceRepoDB reads the source repo path from the Settings table.
func loadSourceRepoDB() string {
	db, err := store.OpenDefault()
	if err != nil {
		return ""
	}
	defer db.Close()

	return db.GetSetting(constants.SettingSourceRepoPath)
}

// findGitRoot walks up from dir looking for a .git directory.
func findGitRoot(dir string) string {
	for {
		if info, err := os.Stat(filepath.Join(dir, ".git")); err == nil && info.IsDir() {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}

		dir = parent
	}
}

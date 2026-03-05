package cmd

import (
	"fmt"
	"os"

	"github.com/user/gitmap/cloner"
	"github.com/user/gitmap/model"
)

// runClone handles the "clone" subcommand.
func runClone(args []string) {
	source, targetDir := parseCloneFlags(args)
	if len(source) == 0 {
		fmt.Fprintln(os.Stderr, "Error: source file is required")
		fmt.Fprintln(os.Stderr, "Usage: gitmap clone <source> [--target-dir <dir>]")
		os.Exit(1)
	}
	executeClone(source, targetDir)
}

// executeClone runs the clone operation and prints the summary.
func executeClone(source, targetDir string) {
	summary, err := cloner.CloneFromFile(source, targetDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Clone error: %v\n", err)
		os.Exit(1)
	}
	printSummary(summary)
}

// printSummary displays clone results to the user.
func printSummary(s model.CloneSummary) {
	fmt.Printf("\nClone complete: %d succeeded, %d failed\n",
		s.Succeeded, s.Failed)
	if s.Failed > 0 {
		printFailures(s)
	}
}

// printFailures lists each failed clone operation.
func printFailures(s model.CloneSummary) {
	fmt.Println("\nFailed clones:")
	for _, e := range s.Errors {
		fmt.Printf("  - %s (%s): %s\n",
			e.Record.RepoName, e.Record.RelativePath, e.Error)
	}
}

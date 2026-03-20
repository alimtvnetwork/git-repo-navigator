// Package cmd implements the CLI commands for gitmap.
package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/user/gitmap/constants"
	"github.com/user/gitmap/release"
)

// runReleasePending handles the 'release-pending' command.
func runReleasePending(args []string) {
	checkHelp("release-pending", args)
	assets, notes, draft, dryRun, verbose, noCommit := parseReleasePendingFlags(args)
	_ = verbose

	err := release.ExecutePending(assets, notes, draft, dryRun, noCommit)
	if err != nil {
		fmt.Fprintf(os.Stderr, constants.ErrBareFmt, err)
		os.Exit(1)
	}
}

// parseReleasePendingFlags parses flags for the release-pending command.
func parseReleasePendingFlags(args []string) (assets, notes string, draft, dryRun, verbose, noCommit bool) {
	fs := flag.NewFlagSet(constants.CmdReleasePending, flag.ExitOnError)
	assetsFlag := fs.String("assets", "", constants.FlagDescAssets)
	notesFlag := fs.String("notes", "", constants.FlagDescNotes)
	draftFlag := fs.Bool("draft", false, constants.FlagDescDraft)
	dryRunFlag := fs.Bool("dry-run", false, constants.FlagDescDryRun)
	verboseFlag := fs.Bool("verbose", false, constants.FlagDescVerbose)
	noCommitFlag := fs.Bool("no-commit", false, constants.FlagDescNoCommit)

	// Register -N as shorthand for --notes.
	fs.StringVar(notesFlag, "N", "", constants.FlagDescNotes)

	fs.Parse(args)

	return *assetsFlag, *notesFlag, *draftFlag, *dryRunFlag, *verboseFlag, *noCommitFlag
}

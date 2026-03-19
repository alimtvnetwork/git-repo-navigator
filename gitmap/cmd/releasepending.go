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
	assets, draft, dryRun, verbose, noCommit := parseReleasePendingFlags(args)
	_ = verbose

	err := release.ExecutePending(assets, draft, dryRun, noCommit)
	if err != nil {
		fmt.Fprintf(os.Stderr, constants.ErrBareFmt, err)
		os.Exit(1)
	}
}

// parseReleasePendingFlags parses flags for the release-pending command.
func parseReleasePendingFlags(args []string) (assets string, draft, dryRun, verbose, noCommit bool) {
	fs := flag.NewFlagSet(constants.CmdReleasePending, flag.ExitOnError)
	assetsFlag := fs.String("assets", "", constants.FlagDescAssets)
	draftFlag := fs.Bool("draft", false, constants.FlagDescDraft)
	dryRunFlag := fs.Bool("dry-run", false, constants.FlagDescDryRun)
	verboseFlag := fs.Bool("verbose", false, constants.FlagDescVerbose)
	noCommitFlag := fs.Bool("no-commit", false, constants.FlagDescNoCommit)
	fs.Parse(args)

	return *assetsFlag, *draftFlag, *dryRunFlag, *verboseFlag, *noCommitFlag
}

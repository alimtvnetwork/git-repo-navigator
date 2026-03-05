// Package cmd implements the CLI commands for gitmap.
package cmd

import (
	"flag"
	"fmt"
	"os"
)

// Run is the main entry point for the CLI.
func Run() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	command := os.Args[1]
	dispatch(command)
}

// dispatch routes to the correct subcommand handler.
func dispatch(command string) {
	if command == "scan" {
		runScan(os.Args[2:])
		return
	}
	if command == "clone" {
		runClone(os.Args[2:])
		return
	}
	if command == "help" {
		printUsage()
		return
	}
	fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
	printUsage()
	os.Exit(1)
}

// printUsage displays help text for all commands.
func printUsage() {
	fmt.Println("Usage: gitmap <command> [flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  scan [dir]          Scan directory for Git repos")
	fmt.Println("  clone <source>      Re-clone from CSV/JSON/text file")
	fmt.Println("  help                Show this help message")
	fmt.Println()
	fmt.Println("Run 'gitmap scan --help' or 'gitmap clone --help' for flag details.")
}

// parseScanFlags parses flags for the scan command.
func parseScanFlags(args []string) (dir, configPath, mode, output, outFile string) {
	fs := flag.NewFlagSet("scan", flag.ExitOnError)
	cfgFlag := fs.String("config", "./data/config.json", "Path to config file")
	modeFlag := fs.String("mode", "", "Clone URL style: https or ssh")
	outputFlag := fs.String("output", "", "Output format: terminal, csv, json")
	outFileFlag := fs.String("out-file", "", "Output file path")
	fs.Parse(args)

	dir = "."
	if fs.NArg() > 0 {
		dir = fs.Arg(0)
	}
	return dir, *cfgFlag, *modeFlag, *outputFlag, *outFileFlag
}

// parseCloneFlags parses flags for the clone command.
func parseCloneFlags(args []string) (source, targetDir string) {
	fs := flag.NewFlagSet("clone", flag.ExitOnError)
	targetFlag := fs.String("target-dir", ".", "Base directory for cloned repos")
	fs.Parse(args)

	source = ""
	if fs.NArg() > 0 {
		source = fs.Arg(0)
	}
	return source, *targetFlag
}

import { useState, useMemo } from "react";
import DocsLayout from "@/components/docs/DocsLayout";
import CommandCard from "@/components/docs/CommandCard";
import SearchBar from "@/components/docs/SearchBar";

const commands = [
  {
    name: "scan", alias: "s", description: "Scan directory for Git repos",
    usage: "gitmap scan [dir] [--output csv|json|terminal] [--mode ssh|https]",
    flags: [
      { flag: "--config <path>", description: "Config file path" },
      { flag: "--mode ssh|https", description: "Clone URL style (default: https)" },
      { flag: "--output csv|json|terminal", description: "Output format (default: terminal)" },
      { flag: "--output-path <dir>", description: "Output directory" },
      { flag: "--github-desktop", description: "Add repos to GitHub Desktop" },
      { flag: "--open", description: "Open output folder after scan" },
      { flag: "--quiet", description: "Suppress clone help section" },
    ],
    examples: [
      { command: "gitmap scan ~/projects", description: "Scan a directory" },
      { command: "gitmap s --output json --mode ssh", description: "JSON output with SSH URLs" },
    ],
  },
  {
    name: "clone", alias: "c", description: "Re-clone repos from structured file",
    usage: "gitmap clone <source|json|csv|text> [--target-dir <dir>] [--safe-pull]",
    flags: [
      { flag: "--target-dir <dir>", description: "Base directory for clones" },
      { flag: "--safe-pull", description: "Pull existing repos with retry + diagnostics" },
      { flag: "--verbose", description: "Write detailed debug log" },
    ],
    examples: [
      { command: "gitmap clone json --target-dir ./projects" },
      { command: "gitmap c csv" },
    ],
  },
  {
    name: "pull", alias: "p", description: "Pull a specific repo by name",
    usage: "gitmap pull <repo-name> [--group <name>] [--all] [--verbose]",
    flags: [
      { flag: "--group <name>", description: "Pull all repos in a group" },
      { flag: "--all", description: "Pull all tracked repos" },
      { flag: "--verbose", description: "Enable verbose logging" },
    ],
  },
  {
    name: "status", alias: "st", description: "Show repo status dashboard",
    usage: "gitmap status [--group <name>] [--all]",
  },
  {
    name: "exec", alias: "x", description: "Run git command across all repos",
    usage: "gitmap exec <git-args...>",
    examples: [{ command: "gitmap exec fetch --prune" }],
  },
  {
    name: "watch", alias: "w", description: "Live-refresh repo status dashboard",
    usage: "gitmap watch [--interval <seconds>] [--group <name>] [--no-fetch] [--json]",
    flags: [
      { flag: "--interval <seconds>", description: "Refresh interval (default: 30, min: 5)" },
      { flag: "--group <name>", description: "Monitor only repos in a group" },
      { flag: "--no-fetch", description: "Skip git fetch" },
      { flag: "--json", description: "Output single snapshot as JSON" },
    ],
  },
  {
    name: "release", alias: "r", description: "Create release branch, tag, and push",
    usage: "gitmap release [version] [--bump major|minor|patch] [--draft] [--dry-run]",
    flags: [
      { flag: "--assets <path>", description: "Attach files to release" },
      { flag: "--commit <sha>", description: "Release from specific commit" },
      { flag: "--branch <name>", description: "Release from branch" },
      { flag: "--bump major|minor|patch", description: "Auto-increment version" },
      { flag: "--draft", description: "Create unpublished draft" },
      { flag: "--dry-run", description: "Preview without executing" },
    ],
  },
  {
    name: "cd", alias: "go", description: "Navigate to a tracked repo directory",
    usage: "gitmap cd <repo-name|repos> [--group <name>] [--pick]",
    examples: [
      { command: "gitmap cd myrepo", description: "Navigate to repo" },
      { command: "gitmap cd repos", description: "Interactive repo picker" },
      { command: "gitmap cd repos --group work", description: "Pick from group" },
    ],
  },
  {
    name: "diff-profiles", alias: "dp", description: "Compare repos across two profiles",
    usage: "gitmap diff-profiles <profileA> <profileB> [--all] [--json]",
  },
  {
    name: "list", alias: "ls", description: "Show all tracked repos with slugs",
    usage: "gitmap list [--group <name>] [--verbose]",
  },
  {
    name: "group", alias: "g", description: "Manage repo groups",
    usage: "gitmap group <create|add|remove|list|show|delete> [args]",
  },
  {
    name: "profile", alias: "pf", description: "Manage database profiles",
    usage: "gitmap profile <create|list|switch|delete|show> [name]",
  },
  {
    name: "latest-branch", alias: "lb", description: "Find most recently updated remote branch",
    usage: "gitmap latest-branch [--top N] [--format json|csv|terminal]",
  },
  {
    name: "changelog", alias: "cl", description: "Show release notes",
    usage: "gitmap changelog [version] [--open] [--source]",
  },
  {
    name: "doctor", alias: undefined, description: "Diagnose PATH, deploy, and version issues",
    usage: "gitmap doctor [--fix-path]",
  },
  {
    name: "update", alias: undefined, description: "Self-update from source repo",
    usage: "gitmap update",
  },
  {
    name: "version", alias: "v", description: "Show version number",
    usage: "gitmap version",
  },
  {
    name: "seo-write", alias: "sw", description: "Auto-commit SEO messages",
    usage: "gitmap seo-write [flags]",
  },
  {
    name: "amend", alias: "am", description: "Rewrite commit author info",
    usage: "gitmap amend [commit-hash] --name <name> --email <email>",
  },
  {
    name: "bookmark", alias: "bk", description: "Save and run bookmarked commands",
    usage: "gitmap bookmark <save|list|run|delete> [args]",
  },
  {
    name: "export", alias: "exp", description: "Export database to file",
    usage: "gitmap export [--json]",
  },
  {
    name: "import", alias: "imp", description: "Import repos from file",
    usage: "gitmap import <file>",
  },
  {
    name: "revert", alias: undefined, description: "Revert to a specific release version",
    usage: "gitmap revert <version>",
  },
  {
    name: "gomod", alias: "gm", description: "Rename Go module path across repo with branch safety",
    usage: "gitmap gomod <new-module-path> [--dry-run] [--no-merge] [--no-tidy] [--verbose]",
    flags: [
      { flag: "--dry-run", description: "Preview changes without modifying files or branches" },
      { flag: "--no-merge", description: "Commit on feature branch but do not merge back" },
      { flag: "--no-tidy", description: "Skip go mod tidy after replacement" },
      { flag: "--verbose", description: "Print each file path as it is modified" },
    ],
    examples: [
      { command: 'gitmap gomod "x/y"', description: "Rename module path to x/y" },
      { command: 'gitmap gomod "github.com/new/name" --dry-run', description: "Preview what would change" },
      { command: 'gitmap gomod "github.com/new/name" --no-merge', description: "Replace but stay on feature branch" },
    ],
  },
];

const CommandsPage = () => {
  const [search, setSearch] = useState("");

  const filtered = useMemo(() => {
    if (!search) return commands;
    const q = search.toLowerCase();
    return commands.filter(
      (c) =>
        c.name.includes(q) ||
        c.alias?.includes(q) ||
        c.description.toLowerCase().includes(q)
    );
  }, [search]);

  return (
    <DocsLayout>
      <h1 className="text-3xl font-mono font-bold mb-2">Command Reference</h1>
      <p className="text-muted-foreground mb-6">
        All available gitmap commands with flags and examples.
      </p>

      <SearchBar value={search} onChange={setSearch} />

      <div className="mt-6 space-y-2">
        {filtered.map((cmd) => (
          <CommandCard key={cmd.name} {...cmd} />
        ))}
        {filtered.length === 0 && (
          <p className="text-center text-muted-foreground py-8 font-mono text-sm">
            No commands matching "{search}"
          </p>
        )}
      </div>
    </DocsLayout>
  );
};

export default CommandsPage;

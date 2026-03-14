# gitmap list

Show all tracked repositories with their slugs and paths.

## Alias

ls

## Usage

    gitmap list [--group <name>] [--verbose]
    gitmap ls go              List only Go projects
    gitmap ls node            List only Node.js projects
    gitmap ls react           List only React projects
    gitmap ls cpp             List only C++ projects
    gitmap ls csharp          List only C# projects
    gitmap ls groups          List all groups

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| --group \<name\> | — | Filter to a specific group |
| --verbose | false | Show full paths and metadata |

## Prerequisites

- Run `gitmap scan` first to populate the database (see scan.md)

## Examples

### Example 1: List all tracked repos

    gitmap list

**Output:**

    my-api       ~/projects/my-api
    web-app      ~/projects/web-app
    3 repos tracked

### Example 2: List only Go projects

    gitmap ls go

**Output:**

    go     github.com/user/my-api
           Path: ~/projects/my-api

### Example 3: List all groups

    gitmap ls groups

**Output:**

    GROUP           REPOS   DESCRIPTION
    backend         3       All backend services

## See Also

- [cd](cd.md) — Navigate to a tracked repo
- [group](group.md) — Manage repo groups
- [scan](scan.md) — Scan directories to populate the database
- [status](status.md) — View repo statuses

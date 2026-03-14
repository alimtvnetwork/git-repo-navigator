# gitmap group

Manage repository groups and activate a group for batch operations.

## Alias

g

## Usage

    gitmap group <create|add|remove|list|show|delete|pull|status|exec|clear> [args]
    gitmap g <name>           Activate a group
    gitmap g                  Show active group
    gitmap g pull             Pull repos in active group
    gitmap g status           Show status for active group
    gitmap g exec <args>      Run git across active group
    gitmap g clear            Clear active group

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| --desc \<text\> | — | Group description (for create) |
| --color \<hex\> | — | Group color (for create) |

## Prerequisites

- Run `gitmap scan` first to populate the database (see scan.md)

## Examples

### Example 1: Activate a group

    gitmap g backend

**Output:**

    Active group set: backend

### Example 2: Pull all repos in active group

    gitmap g pull

**Output:**

    Pulling my-api (main)...
    ✓ my-api is up to date.

### Example 3: Show active group

    gitmap g

**Output:**

    Active group: backend

## See Also

- [list](list.md) — View all tracked repos
- [multi-group](multi-group.md) — Select multiple groups
- [pull](pull.md) — Pull repos by group
- [status](status.md) — View status by group

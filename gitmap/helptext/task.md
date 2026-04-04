# gitmap task

Manage named file-sync watch tasks for one-way folder synchronization.

## Alias

tk

## Usage

    gitmap task <subcommand> [flags]

## Subcommands

| Subcommand | Description |
|------------|-------------|
| create | Create a new sync task |
| list | List all saved tasks |
| run | Start a task's sync loop |
| show | Show details of a task |
| delete | Remove a saved task |

## Flags (create)

| Flag | Default | Description |
|------|---------|-------------|
| --src | (required) | Source directory path |
| --dest | (required) | Destination directory path |

## Flags (run)

| Flag | Default | Description |
|------|---------|-------------|
| --interval | 5 | Sync interval in seconds (minimum 2) |
| --verbose | false | Show detailed sync output |
| --dry-run | false | Preview sync actions without copying |

## Prerequisites

- Source directory must exist

## Examples

### Example 1: Create a sync task

    gitmap task create my-sync --src ./src --dest ./backup

**Output:**

    Task 'my-sync' created.

### Example 2: Run a sync task

    gitmap tk run my-sync --interval 10 --verbose

**Output:**

    Task 'my-sync' running — syncing every 10s (Ctrl+C to stop)
    Synced: main.go
    Synced: config/settings.json
    All files up to date.

### Example 3: List all tasks

    gitmap task list

**Output:**

    Tasks:
      my-sync              ./src → ./backup
      docs-mirror          ./docs → /mnt/share/docs

## See Also

- [watch](watch.md) — Live-refresh dashboard of repo status
- [exec](exec.md) — Run git commands across repos
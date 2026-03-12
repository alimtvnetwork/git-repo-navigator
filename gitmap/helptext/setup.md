# gitmap setup

Interactive first-time configuration wizard that applies global Git settings and installs shell tab-completion.

## Alias

None

## Usage

    gitmap setup [--config <path>] [--dry-run]

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| --config \<path\> | ./data/git-setup.json | Path to git-setup.json config file |
| --dry-run | false | Preview changes without applying them |

## Prerequisites

- Git must be installed

## Examples

### Example 1: Run setup wizard

    gitmap setup

**Output:**

    Applying global Git configuration...
    ✓ 3 settings applied
    ■ Shell Completion — powershell
    Shell completion installed for powershell
    Setup complete.

### Example 2: Dry-run mode

    gitmap setup --dry-run

**Output:**

    [DRY RUN] No changes will be made
    [dry-run] would install powershell completion

## See Also

- [completion](completion.md) — Generate completion scripts manually
- [scan](scan.md) — Scan directories after setup
- [doctor](doctor.md) — Diagnose installation issues

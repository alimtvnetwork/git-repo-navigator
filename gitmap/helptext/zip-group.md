# gitmap zip-group

Manage named collections of files and folders that are automatically
compressed into ZIP archives during a release.

## Alias

z

## Usage

    gitmap zip-group <subcommand> [arguments]

## Subcommands

| Subcommand | Description |
|------------|-------------|
| create     | Create a named zip group |
| add        | Add files or folders to a group |
| remove     | Remove an item from a group |
| list       | List all zip groups |
| show       | Show items in a group |
| delete     | Delete a zip group |
| rename     | Set a custom archive name for a group |

## Flags

| Flag | Description |
|------|-------------|
| --archive \<name\> | Custom output filename (used with create/rename) |

## Examples

### Example 1: Create a group and add items

    gitmap z create docs-bundle
    gitmap z add docs-bundle ./README.md ./CHANGELOG.md ./docs/

### Example 2: Create with custom archive name

    gitmap z create extras --archive extra-files.zip
    gitmap z add extras ./config/ ./scripts/deploy.sh

### Example 3: Use during release

    gitmap release v3.0.0 --zip-group docs-bundle
    gitmap release v3.0.0 -Z ./dist/report.pdf --bundle reports.zip

## See Also

- [release](release.md) — Create a release with zip group assets
- [group](group.md) — Manage repository groups

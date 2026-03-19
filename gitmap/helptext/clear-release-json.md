# gitmap clear-release-json

Remove a specific release metadata JSON file from the `.release/` directory.

## Alias

crj

## Usage

    gitmap clear-release-json <version>

## Flags

| Flag | Description |
|------|-------------|
| `--dry-run` | Preview which file would be removed without deleting it |

## Prerequisites

- A `.release/vX.Y.Z.json` file must exist for the given version.

## Examples

### Example 1: Remove a release JSON file

    gitmap clear-release-json v2.20.0

**Output:**

    Found .release/v2.20.0.json (1.2 KB)
    ✓ Removed .release/v2.20.0.json

### Example 2: Dry-run preview

    gitmap clear-release-json v2.20.0 --dry-run

**Output:**

    [dry-run] Found .release/v2.20.0.json (1.2 KB)
    [dry-run] Would remove .release/v2.20.0.json
    No changes made.

### Example 3: Version not found

    gitmap clear-release-json v9.9.9

**Output:**

    ✗ Error: no release file found for v9.9.9
    Available versions in .release/:
      v2.22.0, v2.21.0, v2.20.0, v2.19.0
    → Use 'gitmap list-releases' to see all stored releases

### Example 4: Clear after orphaned metadata prompt

    gitmap release --bump patch
    # ⚠ Release metadata exists for v2.20.0 but no tag found
    # User decides to clean up manually:
    gitmap crj v2.20.0

**Output:**

    Found .release/v2.20.0.json (1.2 KB)
    ✓ Removed .release/v2.20.0.json
    → You can now re-run 'gitmap release --bump patch'

## See Also

- [release](release.md) — Create a release
- [list-releases](list-releases.md) — Show stored releases
- [db-reset](db-reset.md) — Clear the entire database

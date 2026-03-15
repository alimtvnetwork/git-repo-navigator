# gitmap release-pending

Release all pending versions from two sources: local `release/v*`
branches missing tags, and `.release/v*.json` metadata files where
neither the branch nor the tag exists.

## Alias

rp

## Usage

    gitmap release-pending [--assets <path>] [--draft] [--dry-run] [--verbose]

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| --assets | (none) | Directory or file to attach to each release |
| --draft | false | Mark releases as drafts |
| --dry-run | false | Preview without executing |
| --verbose | false | Detailed output |

## Prerequisites

- Must be inside a Git repository

## Examples

### Example 1: Release pending branches

    gitmap release-pending

**Output:**

    Found 2 pending release branch(es).
    ✓ Created tag v2.18.0
    ✓ Release v2.18.0 complete.

### Example 2: Release from metadata

    gitmap rp

**Output:**

    Found 1 pending release branch(es).
    → Found 1 unreleased version(s) from .release/ metadata
    → Creating release from metadata: v2.19.0 (commit: abc1234)
    ✓ Release v2.19.0 complete.

### Example 3: Dry run

    gitmap rp --dry-run

**Output:**

    [dry-run] Create branch release/v2.19.0 from commit abc1234
    [dry-run] Create tag v2.19.0

## See Also

- [release](release.md) — Create a release
- [release-branch](release-branch.md) — Complete from existing branch
- [changelog](changelog.md) — View release notes

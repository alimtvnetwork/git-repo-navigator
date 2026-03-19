# gitmap list-releases

List release metadata stored in the local database.

## Alias

lr

## Usage

    gitmap list-releases [--json] [--source manual|scan]

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| --json | false | Output as structured JSON |
| --source \<type\> | — | Filter by release source (manual or scan) |

## Prerequisites

- Run `gitmap scan` or `gitmap release` to populate release data (see scan.md, release.md)

## Examples

### Example 1: List all stored releases

    gitmap list-releases

**Output:**

    VERSION   DATE          SOURCE   COMMITS  BRANCH            DRAFT
    v2.22.0   2025-03-10    manual   5        release/v2.22.0   no
    v2.21.0   2025-03-08    manual   3        release/v2.21.0   no
    v2.20.0   2025-02-28    manual   4        release/v2.20.0   no
    v2.19.0   2025-02-20    scan     6        release/v2.19.0   no
    v2.18.0   2025-02-15    scan     2        release/v2.18.0   no
    5 releases found

### Example 2: Filter by source (scan-imported only)

    gitmap lr --source scan

**Output:**

    VERSION   DATE          SOURCE   COMMITS
    v2.19.0   2025-02-20    scan     6
    v2.18.0   2025-02-15    scan     2
    2 releases found (filtered: source=scan)

### Example 3: JSON output

    gitmap lr --json

**Output:**

    [
      {"version":"v2.22.0","date":"2025-03-10","source":"manual","commits":5,"draft":false},
      {"version":"v2.21.0","date":"2025-03-08","source":"manual","commits":3,"draft":false}
    ]

## See Also

- [list-versions](list-versions.md) — List Git release tags
- [changelog](changelog.md) — View release notes
- [release](release.md) — Create a release
- [scan](scan.md) — Scan to import release data

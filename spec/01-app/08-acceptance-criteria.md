# Acceptance Criteria

## Scan Feature

- **Given** a directory with 3 nested Git repos,
  **when** `gitmap scan ./dir` is run,
  **then** all 3 repos appear in terminal output with HTTPS clone URLs
  and correct branch names.

- **Given** `--mode ssh`,
  **then** clone instructions use `git@github.com:…` format.

- **Given** `--output csv`,
  **then** a valid CSV file is written with headers:
  `repoName,httpsUrl,sshUrl,branch,relativePath,absolutePath,cloneInstruction,notes`.

- **Given** a folder with no `.git`,
  **then** it is skipped silently.

- **Given** a repo with no remote,
  **then** URL fields are empty, notes say "no remote configured."

## Clone Feature

- **Given** a valid CSV from scan,
  **when** `gitmap clone ./output/gitmap.csv --target-dir ./restored`,
  **then** all repos are cloned into correct relative paths.

- **Given** a repo that fails to clone,
  **then** it is logged and remaining repos continue.
  Summary shows N succeeded, M failed.

## Config Feature

- **Given** no `--config` flag,
  **then** `./data/config.json` is loaded if it exists.

- **Given** CLI flags that conflict with config,
  **then** CLI flags take precedence.

## Code Quality

- No `if` condition uses negation.
- Every function is 8–15 lines.
- Every file is 100–200 lines.
- Each package has a single clear responsibility.

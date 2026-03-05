# Cloner

## Responsibility

Read a structured file (CSV, JSON, or text) and re-clone repositories,
preserving the original folder hierarchy.

## Behavior

1. Detect file format by extension (`.csv`, `.json`, `.txt`).
2. Parse records from the file.
3. For each record:
   a. Create the relative directory structure under `--target-dir`.
   b. Run `git clone -b <branch> <url> <target-path>`.
4. Log success or failure for each clone operation.
5. Print a summary: N succeeded, M failed.

## Error Handling

- If a clone fails (network, auth, etc.), log the error and continue.
- Do not abort the entire run for a single failure.
- Summary at end lists all failures with reasons.

## Input Formats

| Format | Structure                              |
|--------|----------------------------------------|
| CSV    | Standard CSV with headers              |
| JSON   | Array of `ScanRecord` objects          |
| Text   | One `git clone …` command per line     |

# CLI Interface

## Commands

### `gitmap scan [dir]`

Scan `dir` recursively for Git repositories.  
Default: current working directory.

### `gitmap clone <source-file>`

Re-clone repositories from a CSV, JSON, or text file.

## Global Flags

| Flag              | Description                         | Default              |
|-------------------|-------------------------------------|----------------------|
| `--config <path>` | Path to JSON config file            | `./data/config.json` |

## Scan Flags

| Flag                  | Description                | Default    |
|-----------------------|----------------------------|------------|
| `--mode ssh \| https` | Clone URL style            | `https`    |
| `--output csv \| json \| terminal` | Output format | `terminal` |
| `--out-file <path>`   | File path for output       | auto       |

## Clone Flags

| Flag                   | Description                          | Default |
|------------------------|--------------------------------------|---------|
| `--target-dir <path>`  | Base dir to recreate folder structure | `.`     |

## Examples

```bash
gitmap scan ./projects --mode ssh --output csv
gitmap scan --output json --out-file repos.json
gitmap clone ./output/gitmap.csv --target-dir ./restored
```

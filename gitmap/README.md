# GitMap

> Scan directories for Git repositories, generate clone instructions, and re-clone them anywhere.

## Quick Start

### Build

```powershell
# From the gitmap/ directory:
.\build.ps1

# Skip git pull:
.\build.ps1 -NoPull

# Build only, no deploy:
.\build.ps1 -NoPull -NoDeploy

# Deploy to custom path:
.\build.ps1 -DeployPath "D:\tools"
```

The binary and `data/` config folder are output to `./bin/`. By default, the binary is also copied to the deploy path in `powershell.json` (default: `E:\bin-run`).

### Manual Build

```bash
cd gitmap
go build -o ../bin/gitmap.exe .
```

---

## Usage

### Scan a directory

```bash
# Scan current directory, print to terminal (HTTPS mode)
gitmap scan

# Scan a specific folder with SSH URLs
gitmap scan ./projects --mode ssh

# Output CSV
gitmap scan ./projects --output csv

# Output JSON to a specific folder
gitmap scan ./projects --output json --output-path ./my-exports

# Combine: scan, output both CSV and JSON to default gitmap-output/
gitmap scan ./projects --output csv
gitmap scan ./projects --output json
```

### Output path behavior

| Flag | Behavior |
|------|----------|
| No `--output-path` | Creates `gitmap-output/` in current directory |
| `--output-path ./exports` | Writes to `./exports/` |
| `--out-file report.csv` | Writes to exact file path |

### Clone from a previous scan

```bash
# Clone from CSV
gitmap clone ./gitmap-output/gitmap.csv --target-dir ./restored

# Clone from JSON
gitmap clone ./gitmap-output/gitmap.json --target-dir ./restored
```

---

## Configuration

### `data/config.json`

```json
{
  "defaultMode": "https",
  "defaultOutput": "terminal",
  "outputDir": "./output",
  "excludeDirs": [".cache", "node_modules", "vendor", ".venv"],
  "notes": ""
}
```

CLI flags override config values.

### `powershell.json`

```json
{
  "deployPath": "E:\\bin-run",
  "buildOutput": "./bin",
  "binaryName": "gitmap.exe",
  "goSource": "./gitmap",
  "copyData": true
}
```

---

## CLI Reference

### `gitmap scan [dir]`

| Flag | Description | Default |
|------|-------------|---------|
| `--config <path>` | Config file path | `./data/config.json` |
| `--mode ssh\|https` | Clone URL style | `https` |
| `--output csv\|json\|terminal` | Output format | `terminal` |
| `--output-path <dir>` | Output directory | `./gitmap-output` |
| `--out-file <path>` | Exact output file path | — |

### `gitmap clone <source>`

| Flag | Description | Default |
|------|-------------|---------|
| `--target-dir <path>` | Base clone directory | `.` |

---

## CSV Output Columns

`repoName, httpsUrl, sshUrl, branch, relativePath, absolutePath, cloneInstruction, notes`

---

## Project Structure

```
gitmap/
├── main.go              # Entry point
├── cmd/                  # CLI commands
│   ├── root.go           # Routing & flags
│   ├── scan.go           # Scan command
│   └── clone.go          # Clone command
├── config/               # Config loading
├── scanner/              # Directory walking
├── gitutil/              # Git command wrappers
├── mapper/               # Record building
├── formatter/            # Output (terminal, CSV, JSON)
├── cloner/               # Re-clone logic
├── model/                # Data structures
├── data/                 # Default config
│   └── config.json
├── build.ps1             # PowerShell build script
├── powershell.json       # Build/deploy config
└── go.mod
```

## Specs

See [spec/01-app/](../spec/01-app/) for detailed specifications.

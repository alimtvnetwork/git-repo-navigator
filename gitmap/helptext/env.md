# gitmap env

Manage persistent environment variables and PATH entries across platforms.

## Alias

ev

## Usage

    gitmap env <subcommand> [flags]

## Subcommands

| Subcommand | Description |
|------------|-------------|
| set | Set a persistent environment variable |
| get | Get a managed variable's value |
| delete | Remove a managed variable |
| list | List all managed variables |
| path add | Add a directory to PATH |
| path remove | Remove a directory from PATH |
| path list | List managed PATH entries |

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| --system | false | Target system-level variables (Windows, requires admin) |
| --shell | (auto) | Target shell profile: bash, zsh (Unix only) |
| --verbose | false | Show detailed operation output |
| --dry-run | false | Preview changes without applying |

## Prerequisites

- Windows: setx available (built-in)
- Unix: shell profile (~/.bashrc or ~/.zshrc) writable

## Examples

### Example 1: Set a variable

    gitmap env set GOPATH "/home/user/go"

**Output:**

    Set GOPATH=/home/user/go

### Example 2: Add a directory to PATH

    gitmap ev path add /usr/local/go/bin

**Output:**

    Added to PATH: /usr/local/go/bin

### Example 3: List managed variables

    gitmap env list

**Output:**

    Managed variables:
      GOPATH = /home/user/go
      JAVA_HOME = /usr/lib/jvm/java-17

## See Also

- [install](install.md) — Install developer tools
- [doctor](doctor.md) — Diagnose PATH and version issues
- [setup](setup.md) — Configure Git global settings
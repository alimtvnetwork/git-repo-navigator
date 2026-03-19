# gitmap alias

Assign short names to repositories for quick access from anywhere.

## Alias

a

## Usage

    gitmap alias <subcommand> [arguments]

## Subcommands

| Subcommand | Description |
|------------|-------------|
| set        | Create or update an alias for a repo |
| remove     | Remove an alias |
| list       | List all aliases |
| show       | Show the repo linked to an alias |
| suggest    | Auto-suggest aliases for unaliased repos |

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| --apply | false | Auto-accept all suggestions (with suggest) |

## Examples

### Example 1: Create and use an alias

    gitmap alias set api github/user/api-gateway
    gitmap pull -A api

**Output:**

    ✓ Alias "api" → github/user/api-gateway
    Pulling api-gateway... done

### Example 2: Auto-suggest aliases

    gitmap alias suggest

**Output:**

    api-gateway  → api       Accept? (y/N):
    web-frontend → web       Accept? (y/N):

### Example 3: List all aliases

    gitmap alias list

**Output:**

    api   → github/user/api-gateway
    web   → github/user/web-frontend

## See Also

- [cd](cd.md) — Navigate to a repository
- [exec](exec.md) — Run commands in a repository
- [list](list.md) — List tracked repositories

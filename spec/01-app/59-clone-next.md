# Clone Next

## Command

```
gitmap clone-next <version-arg> [flags]
```

## Alias

```
cn
```

## Responsibility

Clone the next (or a specific) versioned iteration of the current repository
into the parent directory, optionally removing the current version folder and
registering the new clone with GitHub Desktop.

## Terminology

| Term | Meaning |
|------|---------|
| Base name | Repo name without version suffix (e.g., `macro-ahk`) |
| Version suffix | Trailing `-vN` or `-vNN` on the repo name (e.g., `-v11`) |
| Current version | Version suffix of the repo in the current working directory |
| Target version | The version to clone, derived from the argument |

## Version Argument

| Argument | Meaning | Example |
|----------|---------|---------|
| `v++` | Increment current version by 1 | `v11` → `v12` |
| `v<N>` | Jump to exact version N | `v15` → clones `-v15` |

### Edge Case: No Existing Suffix

If the current repo has **no** version suffix (e.g., `macro-ahk`), then:

- `v++` treats the current version as `0` and clones `macro-ahk-v2`
  (first explicit version is v2; the un-suffixed repo is conceptually v1).
- `v<N>` clones `macro-ahk-vN` directly.

If the current repo ends with `-v1`, `v++` clones `-v2`.

## Behavior

1. **Detect current repo**
   a. Read `git config --get remote.origin.url` from the current directory.
   b. Extract the repo owner/org and repo name from the remote URL.
   c. Parse the folder name to identify base name and current version.

2. **Compute target**
   a. Apply the version argument to derive the target version number.
   b. Construct the target repo name: `<base-name>-v<target>`.
   c. Construct the target remote URL using the same host and owner.
   d. Construct the target local path: `<parent-dir>/<target-repo-name>`.

3. **Clone**
   a. Verify the target directory does not already exist.
   b. Run `git clone <target-url> <target-path>`.
   c. Print progress: `Cloning <target-repo-name> into <parent-dir>...`

4. **GitHub Desktop registration**
   a. After a successful clone, register the new repo with GitHub Desktop
      (same mechanism as `clone --github-desktop`).
   b. Print: `✓ Registered <target-repo-name> with GitHub Desktop`

5. **Optional: Remove current version**
   a. After successful clone, prompt the user:
      `Remove current folder <current-folder>? [y/N]`
   b. If confirmed, delete the current directory recursively.
   c. Print: `✓ Removed <current-folder>`

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--delete` | false | Skip prompt and auto-remove current folder after clone |
| `--keep` | false | Skip prompt and keep current folder (no removal question) |
| `--no-desktop` | false | Skip GitHub Desktop registration |
| `--verbose` | false | Enable verbose debug logging |
| `--ssh-key <name>` | (none) | Use a named SSH key for the clone |

If neither `--delete` nor `--keep` is provided, the command prompts interactively.

## URL Construction

The target URL mirrors the current remote's scheme and host:

| Current remote | Target URL |
|----------------|------------|
| `https://github.com/user/repo-v11.git` | `https://github.com/user/repo-v12.git` |
| `git@github.com:user/repo-v11.git` | `git@github.com:user/repo-v12.git` |

## Examples

### Example 1: Increment version

```
D:\wp-work\riseup-asia\macro-ahk-v11> gitmap cn v++

Cloning macro-ahk-v12 into D:\wp-work\riseup-asia...
✓ Cloned macro-ahk-v12
✓ Registered macro-ahk-v12 with GitHub Desktop
Remove current folder macro-ahk-v11? [y/N] n
```

### Example 2: Jump to specific version with auto-delete

```
D:\wp-work\riseup-asia\macro-ahk-v12> gitmap cn v15 --delete

Cloning macro-ahk-v15 into D:\wp-work\riseup-asia...
✓ Cloned macro-ahk-v15
✓ Registered macro-ahk-v15 with GitHub Desktop
✓ Removed macro-ahk-v12
```

### Example 3: No existing suffix

```
D:\wp-work\riseup-asia\macro-ahk> gitmap cn v++

Cloning macro-ahk-v2 into D:\wp-work\riseup-asia...
✓ Cloned macro-ahk-v2
✓ Registered macro-ahk-v2 with GitHub Desktop
Remove current folder macro-ahk? [y/N] y
✓ Removed macro-ahk
```

## Error Handling

| Condition | Behavior |
|-----------|----------|
| Not inside a git repo | Print error, exit 1 |
| Cannot parse remote URL | Print error, exit 1 |
| Target directory already exists | Print error, suggest using `cd`, exit 1 |
| Clone fails (network/auth) | Print error, do not prompt for deletion, exit 1 |
| Deletion fails | Print warning, exit 0 (clone already succeeded) |

## Implementation Scope

| Component | File |
|-----------|------|
| Command handler | `cmd/clonenext.go` |
| Version parser | `clonenext/version.go` |
| Dispatch registration | `cmd/rootcore.go` |
| Constants | `constants/constants_cli.go` |
| Help text | `helptext/clone-next.md` |

## Acceptance Criteria

1. `gitmap cn v++` increments the version suffix and clones into the parent directory.
2. `gitmap cn v<N>` clones the exact version N into the parent directory.
3. Repos without a version suffix are handled correctly (treated as v1).
4. The new clone is registered with GitHub Desktop by default.
5. User is prompted to remove the current folder unless `--delete` or `--keep` is specified.
6. URL scheme (HTTPS or SSH) is preserved from the current remote.
7. All error conditions produce clear messages and correct exit codes.
8. `--verbose` writes a debug log file.
9. Shell completion is provided for the version argument and flags.

## See Also

- [Cloner](05-cloner.md) — Core clone-from-file logic
- [GitHub Desktop](10-github-desktop.md) — Desktop registration
- [SSH Keys](50-ssh.md) — Named SSH key support

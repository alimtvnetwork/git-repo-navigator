# Issue: `gitmap update` fails with "file is being used by another process"

## Root Cause

When `gitmap update` runs from `E:\bin-run\gitmap.exe`, the process holds a file lock on the binary. The update triggers `run.ps1` which tries to `Copy-Item` over the same binary during the deploy step — but the original process hasn't exited yet, so Windows blocks the overwrite.

## Solution (v1.1.1 + v1.1.2)

Two-layer fix:

1. **Copy-and-handoff** (`gitmap/cmd/update.go`):
   - Parent copies itself to `%TEMP%\gitmap-update-<pid>.exe`
   - Launches the copy with `update --from-copy`
   - Parent **exits immediately** (`os.Exit(0)`) to release the file lock
   - The copy then generates and runs a PowerShell script that calls `run.ps1`

2. **Deploy retry** (`run.ps1`):
   - `Copy-Item` wrapped in a retry loop (20 attempts, 500ms delay)
   - Handles race condition where parent hasn't fully released the handle yet

3. **Startup delay** (`gitmap/cmd/update.go`):
   - 1.2s `Start-Sleep` in the generated PowerShell script before calling `run.ps1`

## Iterations

| Version | Change | Result |
|---------|--------|--------|
| v1.1.0 | No handoff — update ran in-process | ❌ File lock error |
| v1.1.1 | Added copy-and-handoff + SSH clone output | ❌ Still locked (timing race) |
| v1.1.2 | Added deploy retry + startup delay | ✅ Should resolve (pending verification) |

## Learnings

- **Windows file locks are held until the process fully terminates** — `cmd.Start()` + `os.Exit(0)` isn't always instant
- **Always add retry logic** for file operations on deployed binaries
- **A delay before rebuild** gives the parent process time to fully release handles
- **Don't assume `os.Exit(0)` releases locks immediately** — the OS may keep the handle briefly

## What NOT to Repeat

- Don't run `run.ps1` (which overwrites the binary) from the same process that holds the lock
- Don't skip the deploy retry — even with the handoff, a small timing window exists
- Always bump the version so the user can confirm the update actually applied

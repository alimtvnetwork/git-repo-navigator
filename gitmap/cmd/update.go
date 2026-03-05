package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/user/gitmap/constants"
)

// runUpdate handles the "update" subcommand.
func runUpdate() {
	repoPath := constants.RepoPath
	if len(repoPath) == 0 {
		fmt.Fprintln(os.Stderr, constants.ErrNoRepoPath)
		os.Exit(1)
	}
	fmt.Printf(constants.MsgUpdateStarting)
	fmt.Printf(constants.MsgUpdateRepoPath, repoPath)
	executeUpdate(repoPath)
}

// executeUpdate writes a temp PS1 script and runs it.
func executeUpdate(repoPath string) {
	scriptPath := writeUpdateScript(repoPath)
	defer os.Remove(scriptPath)

	runUpdateScript(scriptPath)
}

// writeUpdateScript creates a temporary PowerShell script for self-update.
// Writes with UTF-8 BOM so PowerShell correctly handles Unicode characters.
func writeUpdateScript(repoPath string) string {
	runPS1 := filepath.Join(repoPath, "run.ps1")
	script := buildUpdateScript(repoPath, runPS1)
	tmpFile := filepath.Join(os.TempDir(), "gitmap-update.ps1")

	// UTF-8 BOM prefix for PowerShell compatibility
	bom := []byte{0xEF, 0xBB, 0xBF}
	content := append(bom, []byte(script)...)
	os.WriteFile(tmpFile, content, constants.DirPermission)

	return tmpFile
}

// buildUpdateScript generates the PowerShell script content.
// Uses ASCII-safe characters to avoid encoding issues.
func buildUpdateScript(repoPath, runPS1 string) string {
	return fmt.Sprintf(`# gitmap self-update script (auto-generated)
Set-Location "%s"
Write-Host ""
Write-Host "  Pulling and rebuilding gitmap..." -ForegroundColor Cyan
Write-Host ""
& "%s"
Write-Host ""
$newBinary = Join-Path "%s" "bin\gitmap.exe"
if (Test-Path $newBinary) {
    $version = & $newBinary help 2>&1 | Select-String -Pattern "v\d+\.\d+\.\d+" | ForEach-Object { $_.Matches[0].Value }
    Write-Host "  [OK] Updated to gitmap $version" -ForegroundColor Green
} else {
    Write-Host "  [OK] Update complete" -ForegroundColor Green
}
Write-Host ""
`, repoPath, runPS1, repoPath)
}

// runUpdateScript executes the PowerShell script with output piped to terminal.
func runUpdateScript(scriptPath string) {
	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass",
		"-NoProfile", "-File", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, constants.ErrUpdateFailed, err)
		os.Exit(1)
	}
}

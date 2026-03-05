// Package mapper converts raw scan data into ScanRecord structs.
package mapper

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/user/gitmap/gitutil"
	"github.com/user/gitmap/model"
	"github.com/user/gitmap/scanner"
)

// BuildRecords converts a list of RepoInfo into ScanRecords.
func BuildRecords(repos []scanner.RepoInfo, mode, defaultNote string) []model.ScanRecord {
	records := make([]model.ScanRecord, 0, len(repos))
	for _, repo := range repos {
		rec := buildOneRecord(repo, mode, defaultNote)
		records = append(records, rec)
	}
	return records
}

// buildOneRecord creates a single ScanRecord from a RepoInfo.
func buildOneRecord(repo scanner.RepoInfo, mode, note string) model.ScanRecord {
	remoteURL, _ := gitutil.RemoteURL(repo.AbsolutePath)
	branch, _ := gitutil.CurrentBranch(repo.AbsolutePath)
	httpsURL := toHTTPS(remoteURL)
	sshURL := toSSH(remoteURL)
	cloneURL := selectCloneURL(httpsURL, sshURL, mode)
	repoName := extractRepoName(remoteURL)
	noteText := buildNote(remoteURL, note)
	instruction := buildInstruction(cloneURL, branch, repo.RelativePath)

	return model.ScanRecord{
		ID: uuid.New().String(), RepoName: repoName,
		HTTPSUrl: httpsURL, SSHUrl: sshURL, Branch: branch,
		RelativePath: repo.RelativePath, AbsolutePath: repo.AbsolutePath,
		CloneInstruction: instruction, Notes: noteText,
	}
}

// toHTTPS converts a remote URL to HTTPS format.
func toHTTPS(raw string) string {
	if strings.HasPrefix(raw, "https://") {
		return raw
	}
	if strings.HasPrefix(raw, "git@") {
		host, path := splitSSH(raw)
		return fmt.Sprintf("https://%s/%s", host, path)
	}
	return raw
}

// toSSH converts a remote URL to SSH format.
func toSSH(raw string) string {
	if strings.HasPrefix(raw, "git@") {
		return raw
	}
	if strings.HasPrefix(raw, "https://") {
		trimmed := strings.TrimPrefix(raw, "https://")
		parts := strings.SplitN(trimmed, "/", 2)
		if len(parts) == 2 {
			return fmt.Sprintf("git@%s:%s", parts[0], parts[1])
		}
	}
	return raw
}

// splitSSH splits a git@host:path URL into host and path.
func splitSSH(raw string) (string, string) {
	trimmed := strings.TrimPrefix(raw, "git@")
	parts := strings.SplitN(trimmed, ":", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return trimmed, ""
}

// selectCloneURL picks HTTPS or SSH URL based on mode.
func selectCloneURL(httpsURL, sshURL, mode string) string {
	if mode == "ssh" {
		return sshURL
	}
	return httpsURL
}

// extractRepoName derives the repository name from a remote URL.
func extractRepoName(raw string) string {
	if len(raw) == 0 {
		return "unknown"
	}
	base := filepath.Base(raw)
	return strings.TrimSuffix(base, ".git")
}

// buildNote generates the notes field for a record.
func buildNote(remoteURL, defaultNote string) string {
	if len(remoteURL) == 0 {
		return "no remote configured"
	}
	return defaultNote
}

// buildInstruction creates the full git clone command string.
func buildInstruction(url, branch, relPath string) string {
	if len(url) == 0 {
		return ""
	}
	return fmt.Sprintf("git clone -b %s %s %s", branch, url, relPath)
}

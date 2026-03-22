package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/user/gitmap/constants"
	"github.com/user/gitmap/store"
)

// runSSHGenerate generates a new SSH key pair.
func runSSHGenerate(args []string) {
	name := flagValue(args, constants.FlagSSHName, constants.FlagSSHNameS, constants.DefaultSSHKeyName)
	keyPath := flagValue(args, constants.FlagSSHPath, constants.FlagSSHPathS, defaultSSHKeyPath(name))
	email := flagValue(args, constants.FlagSSHEmail, constants.FlagSSHEmailS, "")
	force := hasAnyFlag(args, constants.FlagSSHForce, constants.FlagSSHForceS)

	if err := validateSSHKeygen(); err != nil {
		fmt.Fprint(os.Stderr, constants.ErrSSHKeygenMissing)
		os.Exit(1)
	}

	if len(email) == 0 {
		email = resolveGitEmail()
	}
	if len(email) == 0 {
		fmt.Fprint(os.Stderr, constants.ErrSSHEmailResolve)
		os.Exit(1)
	}

	keyPath = expandHome(keyPath)

	db := openDB()
	defer db.Close()

	if db.SSHKeyExists(name) && !force {
		if !handleExistingKey(db, name, &keyPath) {
			return
		}
	}

	generateAndStore(db, name, keyPath, email)
}

// handleExistingKey prompts the user when a key already exists.
// Returns true if generation should proceed, false to cancel.
// May update keyPath if user chooses a new path.
func handleExistingKey(db *store.DB, name string, keyPath *string) bool {
	existing, _ := db.FindSSHKeyByName(name)
	fmt.Fprintf(os.Stdout, constants.MsgSSHExists, name, existing.PrivatePath)
	fmt.Fprintf(os.Stdout, constants.MsgSSHExistsFP, existing.Fingerprint)
	fmt.Fprint(os.Stdout, constants.MsgSSHPromptAction)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToUpper(input))

	if input == "R" {
		removeKeyFiles(existing.PrivatePath)
		*keyPath = existing.PrivatePath

		return true
	}
	if input == "N" {
		fmt.Fprint(os.Stdout, constants.MsgSSHNewPathPrompt)
		newPath, _ := reader.ReadString('\n')
		*keyPath = expandHome(strings.TrimSpace(newPath))

		return true
	}

	return false
}

// generateAndStore runs ssh-keygen and stores the result in the database.
func generateAndStore(db *store.DB, name, keyPath, email string) {
	if err := ensureDir(filepath.Dir(keyPath)); err != nil {
		fmt.Fprintf(os.Stderr, constants.ErrSSHKeygen, err)
		os.Exit(1)
	}

	cmd := exec.Command(constants.SSHKeygenBin,
		"-t", constants.SSHKeyType,
		"-b", constants.SSHKeyBits,
		"-C", email,
		"-f", keyPath,
		"-N", "")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, constants.ErrSSHKeygen, err)
		os.Exit(1)
	}

	pubKey, err := os.ReadFile(keyPath + ".pub")
	if err != nil {
		fmt.Fprintf(os.Stderr, constants.ErrSSHReadPub, err)
		os.Exit(1)
	}

	fingerprint := readFingerprint(keyPath)

	if db.SSHKeyExists(name) {
		_ = db.UpdateSSHKey(name, keyPath, string(pubKey), fingerprint, email)
	} else {
		_, _ = db.InsertSSHKey(name, keyPath, string(pubKey), fingerprint, email)
	}

	fmt.Fprintf(os.Stdout, constants.MsgSSHGenerated, name)
	fmt.Fprintf(os.Stdout, constants.MsgSSHPath, keyPath)
	fmt.Fprintf(os.Stdout, constants.MsgSSHFingerprint, fingerprint)
	fmt.Fprint(os.Stdout, constants.MsgSSHPubLabel)
	fmt.Fprintf(os.Stdout, "  %s\n", strings.TrimSpace(string(pubKey)))
	fmt.Fprint(os.Stdout, constants.MsgSSHCopyHint)

	updateSSHConfig(db)
}

// validateSSHKeygen checks if ssh-keygen is available on PATH.
func validateSSHKeygen() error {
	_, err := exec.LookPath(constants.SSHKeygenBin)

	return err
}

// resolveGitEmail reads the global Git email config.
func resolveGitEmail() string {
	out, err := exec.Command("git", "config", "--global", "user.email").Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(out))
}

// readFingerprint reads the SHA256 fingerprint of a key file.
func readFingerprint(keyPath string) string {
	out, err := exec.Command(constants.SSHKeygenBin, "-lf", keyPath+".pub").Output()
	if err != nil {
		return "unknown"
	}

	parts := strings.Fields(string(out))
	if len(parts) >= 2 {
		return parts[1]
	}

	return "unknown"
}

// removeKeyFiles deletes private and public key files.
func removeKeyFiles(privatePath string) {
	_ = os.Remove(privatePath)
	_ = os.Remove(privatePath + ".pub")
}

// defaultSSHKeyPath returns the default key path based on name.
func defaultSSHKeyPath(name string) string {
	home, _ := os.UserHomeDir()
	if name == constants.DefaultSSHKeyName {
		return filepath.Join(home, ".ssh", "id_rsa")
	}

	return filepath.Join(home, ".ssh", "id_rsa_"+name)
}

// expandHome expands ~ to the user's home directory.
func expandHome(path string) string {
	if strings.HasPrefix(path, "~") {
		home, _ := os.UserHomeDir()
		path = filepath.Join(home, path[1:])
	}

	return path
}

// ensureDir creates a directory if it doesn't exist.
func ensureDir(dir string) error {
	return os.MkdirAll(dir, 0700)
}

// hasAnyFlag checks if any of the given flag names appear in args.
func hasAnyFlag(args []string, flags ...string) bool {
	for _, arg := range args {
		for _, f := range flags {
			if arg == f {
				return true
			}
		}
	}

	return false
}

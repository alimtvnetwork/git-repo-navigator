package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/user/gitmap/constants"
)

// runSSHCat displays the public key for a named SSH key.
func runSSHCat(args []string) {
	name := flagValue(args, constants.FlagSSHName, constants.FlagSSHNameS, constants.DefaultSSHKeyName)

	db := openDB()
	defer db.Close()

	key, err := db.FindSSHKeyByName(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, constants.ErrSSHNotFound, name)
		printAvailableKeys(db)
		os.Exit(1)
	}

	fmt.Print(strings.TrimSpace(key.PublicKey))
	fmt.Println()
}

// printAvailableKeys prints available SSH key names to stderr.
func printAvailableKeys(db interface{ SSHKeyNames() ([]string, error) }) {
	names, err := db.(interface{ SSHKeyNames() ([]string, error) }).SSHKeyNames()
	if err != nil || len(names) == 0 {
		return
	}

	fmt.Fprintf(os.Stderr, constants.ErrSSHAvailable, strings.Join(names, ", "))
}

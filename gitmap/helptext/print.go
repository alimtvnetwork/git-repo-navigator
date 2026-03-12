package helptext

import (
	"fmt"
	"os"
)

// Print reads and prints the help file for the given command.
func Print(command string) {
	data, err := Files.ReadFile(command + ".md")
	if err != nil {
		fmt.Fprintf(os.Stderr, "No help available for '%s'\n", command)
		os.Exit(1)
	}
	fmt.Print(string(data))
}

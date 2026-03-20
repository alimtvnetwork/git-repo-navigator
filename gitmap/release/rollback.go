package release

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/user/gitmap/constants"
)

// Rollback deletes a local branch and tag after a failed push.
func Rollback(branchName, tag, originalBranch string) {
	fmt.Fprint(os.Stderr, constants.MsgRollbackStart)

	switchBack(originalBranch)
	deleteLocalBranch(branchName)
	deleteLocalTag(tag)

	fmt.Fprint(os.Stderr, constants.MsgRollbackDone)
}

// switchBack returns to the original branch before deleting the release branch.
func switchBack(branch string) {
	if len(branch) == 0 {
		return
	}

	err := CheckoutBranch(branch)
	if err != nil {
		fmt.Fprintf(os.Stderr, constants.MsgRollbackWarn, "checkout "+branch, err)
	}
}

// deleteLocalBranch force-deletes a local branch.
func deleteLocalBranch(branchName string) {
	if len(branchName) == 0 {
		return
	}

	cmd := exec.Command(constants.GitBin, constants.GitBranch, "-D", branchName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, constants.MsgRollbackWarn, "delete branch "+branchName, err)

		return
	}

	fmt.Fprintf(os.Stderr, constants.MsgRollbackBranch, branchName)
}

// deleteLocalTag deletes a local tag.
func deleteLocalTag(tag string) {
	if len(tag) == 0 {
		return
	}

	cmd := exec.Command(constants.GitBin, constants.GitTag, "-d", tag)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, constants.MsgRollbackWarn, "delete tag "+tag, err)

		return
	}

	fmt.Fprintf(os.Stderr, constants.MsgRollbackTag, tag)
}

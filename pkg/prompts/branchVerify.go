package prompts

import (
	"github.com/manifoldco/promptui"
)

// IsCorrectBranch asks user if they are on the correct branch.
func IsCorrectBranch(branch string) (bool, error) {

	prompt := promptui.Prompt{
		Label:     "You are on branch " + branch + ". Is this correct?",
		IsConfirm: true,
	}

	_, err := prompt.Run()

	if err != nil {
		return false, err
	}

	return true, nil
}

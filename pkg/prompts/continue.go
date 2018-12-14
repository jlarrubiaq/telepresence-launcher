package prompts

import (
	"github.com/manifoldco/promptui"
)

// Continue asks user if they would like to proceed.
func Continue() (bool, error) {

	prompt := promptui.Prompt{
		Label:     "Continue",
		IsConfirm: true,
		Default:   "Y",
	}

	_, err := prompt.Run()

	if err != nil {
		return false, err
	}

	return true, nil
}

package prompts

import (
	"errors"

	"github.com/manifoldco/promptui"
)

// RepoOwner prompts for the repo owner with default provided.
func RepoOwner(defaultval string) (string, error) {

	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("Username must have more than 3 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Github repo owner ",
		Validate: validate,
		Default:  defaultval,
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, err
}

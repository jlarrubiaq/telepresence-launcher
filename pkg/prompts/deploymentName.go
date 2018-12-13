package prompts

import (
	"github.com/manifoldco/promptui"
)

// DeploymentName prompts for the repo owner with default provided.
func DeploymentName(defaultval []string) (string, error) {

	prompt := promptui.Select{
		Label: "Deployment you wish to work on ",
		Items: defaultval,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, err
}

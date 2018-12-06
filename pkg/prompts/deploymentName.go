package prompts

import (
	"github.com/manifoldco/promptui"
)

// DeploymentName prompts for the repo owner with default provided.
func DeploymentName(defaultval map[string]interface{}) (string, error) {
	list := []string{}

	for k := range defaultval {
		list = append(list, k)
	}

	prompt := promptui.Select{
		Label: "Deployment you wish to work on ",
		Items: list,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, err
}

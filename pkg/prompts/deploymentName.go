package prompts

import (
	"github.com/manifoldco/promptui"
)

// DeploymentName prompts for the user to choose which deployment they want to swap with a telepresence instance.
// the list of deployments provided comes directly from .tl.yaml
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

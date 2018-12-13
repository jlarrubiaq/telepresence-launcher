package prompts

import (
	"github.com/manifoldco/promptui"
)

// LaunchMethod prompts to choose the launchmethod
func LaunchMethod(methods []string) (string, error) {

	if len(methods) > 1 {

		prompt := promptui.Select{
			Label: "Which launch method would you like to use? ",
			Items: methods,
		}

		_, result, err := prompt.Run()

		if err != nil {
			return "", err
		}

		return result, err
	}

	return methods[0], nil

}

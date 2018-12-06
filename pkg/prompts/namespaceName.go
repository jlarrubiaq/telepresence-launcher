package prompts

import (
	"strings"

	"github.com/manifoldco/promptui"
)

// NamespaceName prompts for the repo owner with default provided.
func NamespaceName(defaultval []string, repo string) (string, error) {

	filteredList := []string{}
	for _, item := range defaultval {
		if strings.Contains(item, repo) {
			filteredList = append(filteredList, item)
		}
	}

	prompt := promptui.Select{
		Label: "Select the PR you want to work on ",
		Items: filteredList,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, err
}

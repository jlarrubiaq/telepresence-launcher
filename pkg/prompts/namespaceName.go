package prompts

import (
	"strings"

	"github.com/manifoldco/promptui"
)

// NamespaceName prompts to select the namespace you want to use.
// It will attempt to find a namespace where the name of the namespace contains
// the name of the repo found in .tl.yaml. If no match is found, it will return a list of all namespaces
// that contain a deployment matching the repo in .tl.yaml.
func NamespaceName(defaultval []string, repo string) (string, error) {
	foundResult := false
	filteredList := []string{}

	for _, item := range defaultval {
		if strings.Contains(item, repo) {
			filteredList = append(filteredList, item)
			foundResult = true
		}
	}

	if foundResult == false {
		filteredList = defaultval
	}

	prompt := promptui.Select{
		Label: "Select the Namespace you want to use. ",
		Items: filteredList,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, err
}

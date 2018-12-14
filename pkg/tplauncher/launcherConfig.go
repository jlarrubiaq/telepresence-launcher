package tplauncher

import (
	"errors"

	"github.com/ghodss/yaml"
	"github.com/mitchellh/mapstructure"
)

// LauncherConfig is the outer shell of the config file.
type LauncherConfig struct {
	Repo          string `json:"repo"`
	LabelSelector string `json:"labelSelector"`
	Deployments   map[string]struct {
		LaunchMethods []map[string]interface{} `json:"launchMethods"`
	} `json:"deployments"`
}

// NewConfig returns a new config struct based on data provided
func NewConfig(configData []byte) (LauncherConfig, error) {
	var config LauncherConfig
	err := yaml.Unmarshal(configData, &config)

	return config, err
}

// GetDeploymentPartials returns a list of deployments from config.
func (c LauncherConfig) GetDeploymentPartials() []string {
	var list []string
	for name := range c.Deployments {
		list = append(list, name)
	}
	return list
}

// GetAvailableLaunchMethods returns a strign slice of the launch methods available to this deployment
func (c LauncherConfig) GetAvailableLaunchMethods(deploymentPartial string) []string {
	var list []string
	for _, value := range c.Deployments[deploymentPartial].LaunchMethods {
		list = append(list, value["method"].(string))
	}
	return list
}

// GetMethodData returns a type of LaunchMethod based on the deployment name and the method selected by the user.
func (c LauncherConfig) GetMethodData(deploymentPartial string, method string) (LaunchMethod, error) {

	// for each of the methods described in the config
	for _, value := range c.Deployments[deploymentPartial].LaunchMethods {
		if value["method"].(string) == method {
			return decode(value, value["method"].(string))
		}
	}

	return ContainerMethod{}, errors.New("no accepted launch method found in config")
}

// Helper function which uses the mapstructure package to convert the launch method data into a defined Type.
func decode(data interface{}, launchMethod string) (LaunchMethod, error) {
	decoderconfig := mapstructure.DecoderConfig{
		TagName: "json",
	}

	switch launchMethod {
	case "ContainerMethod":
		result := ContainerMethod{}
		decoderconfig.Result = &result
		methodDecoder, _ := mapstructure.NewDecoder(&decoderconfig)
		err := methodDecoder.Decode(data)
		return result, err
	}

	return ContainerMethod{}, errors.New("no accepted launch method could be decoded")
}

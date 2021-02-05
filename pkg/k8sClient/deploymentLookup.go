package k8sClient

import (
	"errors"
	"regexp"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// for GCP support
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// GetDeployment : given a namespace and a slice of deployment partial strings, return options
func (k KubeClient) GetDeployment(namespace string, deploymentPartial string) (string, error) {

	options := metav1.ListOptions{}

	deployments, err := k.clientSet.AppsV1().Deployments(namespace).List(options)

	if err != nil {
		return "", err
	}

	regex := deploymentPartial + "$"
	for _, d := range deployments.Items {
		if match, _ := regexp.MatchString(regex, d.Name); match {
			return d.Name, nil
		}
	}

	return "", errors.New("no deployment match found")

}

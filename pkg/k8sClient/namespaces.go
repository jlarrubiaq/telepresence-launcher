package k8sClient

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// for GCP support
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

//ListNamespaces lists the currently active namespaces.
func (k KubeClient) ListNamespaces(labelSelector string) ([]string, error) {

	options := metav1.ListOptions{
		LabelSelector: labelSelector,
	}

	deployments, err := k.clientSet.AppsV1().Deployments("").List(options)

	if err != nil {
		return []string{}, err
	}

	list := []string{}
	for _, d := range deployments.Items {
		list = append(list, d.Namespace)
	}

	return list, nil

}

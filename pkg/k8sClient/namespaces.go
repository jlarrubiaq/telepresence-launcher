package k8sClient

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// for GCP support
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

//ListNamespaces lists the currently active namespaces.
func (k KubeClient) ListNamespaces() ([]string, error) {

	options := metav1.ListOptions{
		LabelSelector: "app=mwg-singularity",
	}

	deployments, err := k.clientSet.ExtensionsV1beta1().Deployments("").List(options)

	if err != nil {
		return []string{}, err
	}

	list := []string{}
	for _, d := range deployments.Items {
		list = append(list, d.Namespace)
	}

	return list, nil

}

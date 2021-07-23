package k8sClient

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// for GCP support
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

//ListNamespaces lists the currently active namespaces.
func (k KubeClient) ListNamespaces(labelSelector string) ([]string, error) {

	list := []string{}
	// Used to avoid duplicates.
	listMap := make(map[string]bool)
	options := metav1.ListOptions{
		LabelSelector: labelSelector,
	}

	// Try to find the app among Deployments.
	deployments, err := k.clientSet.AppsV1().Deployments("").List(options)

	if err != nil {
		return []string{}, err
	}

	// If there is no luck, try StatefulSets.
	if len(deployments.Items) > 0 {
		for _, d := range deployments.Items {
			if _, ok := listMap[d.Namespace]; !ok {
				listMap[d.Namespace] = true
				list = append(list, d.Namespace)
			}

		}
	} else {
		statefulSets, err := k.clientSet.AppsV1().StatefulSets("").List(options)

		if err != nil {
			return []string{}, err
		}

		for _, s := range statefulSets.Items {
			if _, ok := listMap[s.Namespace]; !ok {
				listMap[s.Namespace] = true
				list = append(list, s.Namespace)
			}
		}
	}

	return list, nil
}

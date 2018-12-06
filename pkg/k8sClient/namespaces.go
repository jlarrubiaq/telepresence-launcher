package k8sClient

import (
	"flag"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// for GCP support
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

//ListNamespaces lists the currently active namespaces.
func ListNamespaces() ([]string, error) {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	options := metav1.ListOptions{
		LabelSelector: "app=mwg-singularity",
	}

	deployments, err := clientset.ExtensionsV1beta1().Deployments("").List(options)

	if err != nil {
		return []string{}, err
	}

	list := []string{}
	for _, d := range deployments.Items {
		list = append(list, d.Namespace)
	}

	return list, nil

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

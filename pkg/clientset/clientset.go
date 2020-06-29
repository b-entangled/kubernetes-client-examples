package clientset

import (
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
)

func GetNewClientset(config *restclient.Config) (*kubernetes.Clientset, error) {
	clientset, err := kubernetes.NewForConfig(config)
	return clientset, err
}


package clientset

import (
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
)

func GetNewClientset(config *restclient.Config) (*kubernetes.Clientset, error) {
	clientset, err := kubernetes.NewForConfig(config)
	return clientset, err
}

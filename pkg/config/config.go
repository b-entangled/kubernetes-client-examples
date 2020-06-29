package config

import (
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetNewConfig(kubeconfig string) (*restclient.Config, error) {
	if kubeconfig == "" {
		return restclient.InClusterConfig()
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	return config, err
}

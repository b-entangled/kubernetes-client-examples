package main

import (
	"log"
	"os"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/b-entangled/kubernetes-client-examples/pkg/clientset"
	"github.com/b-entangled/kubernetes-client-examples/pkg/config"
	"github.com/b-entangled/kubernetes-client-examples/pkg/informers"
	"github.com/b-entangled/kubernetes-client-examples/pkg/resources/configmaps"
)

func main() {
	kubeconfig := os.Getenv("KUBECONFIG")
	config, err := config.GetNewConfig(kubeconfig)
	if err != nil {
		panic("Kubeconfig Error")
	}
	clientset, err := clientset.GetNewClientset(config)
	if err != nil {
		panic("Clientset Error")
	}
	// Create Shared Informers for all Namespace
	informers := informers.NewSharedInformerFactoryWithNamespace(clientset, 0, "")


	// addFunc called when new configmap is added to api-server
	addFunc := func(obj interface{}) {
		mObj := obj.(metav1.Object)
		log.Printf("New ConfigMap Added to Store: %s\n", mObj.GetName())
		
	}

	// addFunc called when existing configmap is deleted from api-server
	deleteFunc := func(obj interface{}) {
		mObj := obj.(metav1.Object)
		log.Printf("Old ConfigMap Deleted from Store: %s\n", mObj.GetName())
	}

	// addFunc called when configmap is updated in api-server
	updateFunc := func(oldObj interface{}, newObj interface{}) {
		mObj := newObj.(metav1.Object)
		log.Printf("Old ConfigMap Updated from Store: %s\n", mObj.GetName())
	}

	configmaps.AddConfigMapsEventInformer(informers, addFunc, deleteFunc, updateFunc)
}

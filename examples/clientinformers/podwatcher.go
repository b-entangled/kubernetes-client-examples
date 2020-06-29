package main

import (
	"fmt"
	"os"
	"k8s.io/client-go/tools/cache"
	apiv1 "k8s.io/api/core/v1"
	"github.com/b-entangled/kubernetes-client-examples/pkg/clientset"
	"github.com/b-entangled/kubernetes-client-examples/pkg/config"
	"k8s.io/client-go/informers/core/v1"
	mv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	informer := v1.NewPodInformer(clientset, "", 0, nil)
	stopper := make(chan struct{})
	defer close(stopper)
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mObj := obj.(mv1.Object)
			fmt.Printf("New Pod Added to Store: %s\n", mObj.GetName())
			pod := obj.(*apiv1.Pod)
			fmt.Printf("Image : %+v\n\n", pod.Spec.Containers[0].Image)
		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(mv1.Object)
			fmt.Printf("Old Pod Deleted from Store: %s", mObj.GetName())
		},
	})

	informer.Run(stopper)
}
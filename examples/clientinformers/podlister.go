package main

import (
	"log"
	"os"
	"time"
	"github.com/b-entangled/kubernetes-client-examples/pkg/clientset"
	"github.com/b-entangled/kubernetes-client-examples/pkg/config"
	"k8s.io/apimachinery/pkg/labels"
	"github.com/b-entangled/kubernetes-client-examples/pkg/informers"
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
	log.Println("Shared Informers for Pods in default Namespace")
	informers := informers.NewSharedInformerFactory(clientset, 0)
	podInformer := informers.SharedInformerFactory.Core().V1().Pods()

	// Create Pod Listerfor all Namespace
	lister := podInformer.Lister().Pods("")
	stopper := make(chan struct{})
	defer close(stopper)
	informers.SharedInformerFactory.Start(stopper)
	time.Sleep(1 * time.Second)
	pods, err := lister.List(labels.Everything())
	
	for _, p := range pods{
		log.Printf("%+v\n", p.Name)
	}
	
	log.Println("Completed Listing Pod in all Namespace")
}
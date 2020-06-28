package informers

import (
	"time"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/informers"
)

type SharedInformer struct {
	SharedInformerFactory informers.SharedInformerFactory
}

func NewSharedInformerFactory(clientset *kubernetes.Clientset, resync time.Duration) *SharedInformer {
	factory := informers.NewSharedInformerFactory(clientset, resync)
	return &SharedInformer{SharedInformerFactory: factory}
}

func NewSharedInformerFactoryWithNamespace(clientset *kubernetes.Clientset, resync time.Duration, namespace string) *SharedInformer {
	factory := informers.NewSharedInformerFactoryWithOptions(clientset, resync, informers.WithNamespace(namespace))
	return &SharedInformer{SharedInformerFactory: factory}
}
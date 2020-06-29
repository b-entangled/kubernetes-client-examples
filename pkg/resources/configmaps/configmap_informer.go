package configmaps

import (
	"github.com/b-entangled/kubernetes-client-examples/pkg/informers"
	"k8s.io/client-go/tools/cache"
	
)

func AddConfigMapsEventInformer(factory *informers.SharedInformer, addFunc func(obj interface{}),
deleteFunc func(obj interface{}), updateFunc func(oldObj interface{}, newObj interface{})){
	
	informer := factory.SharedInformerFactory.Core().V1().ConfigMaps().Informer()
	stopper := make(chan struct{})
	defer close(stopper)
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: addFunc,
		DeleteFunc: deleteFunc,
		UpdateFunc: updateFunc,
	})

	informer.Run(stopper)
}
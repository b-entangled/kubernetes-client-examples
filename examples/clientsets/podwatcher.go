package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "~/.kube/config", "(optional) absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	podsClientAll := clientset.CoreV1().Pods(apiv1.NamespaceAll)
	// Watch Pods for Any Changes
	fmt.Printf("Watching Pods in %q namespace:\n", apiv1.NamespaceAll)
	watcher, err := podsClientAll.Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	cha := watcher.ResultChan()
	for event := range cha {
        pod, ok := event.Object.(*apiv1.Pod)
        if !ok {
            log.Fatal("unexpected type")
		}
		log.Println(event.Type)
        log.Println(pod.GetName())
    }
}
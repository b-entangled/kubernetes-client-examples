package main

import (
	"context"
	"flag"
	"fmt"

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

	// Create Clientset for Pods on default Namespace
	podsClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)
	// List Pods
	fmt.Printf("Listing Pods in namespace %q:\n", apiv1.NamespaceDefault)

	list, err := podsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of Pods in Default Namespace is : ", len(list.Items))
	for _, d := range list.Items {
		fmt.Printf(" * %s \n", d.Name)
	}

	fmt.Println("Listed Pods on default Namespace.")


		// Create Clientset for Pods on all Namespace
		podsClient = clientset.CoreV1().Pods(apiv1.NamespaceAll)
		// List Pods
		fmt.Printf("Listing Pods in namespace %q:\n", apiv1.NamespaceAll)
		list, err = podsClient.List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err)
		}

		fmt.Println("Number of Pods in all Namespace is : ", len(list.Items))

		for _, d := range list.Items {
			fmt.Printf(" * %s \n", d.Name)
		}
	
		fmt.Println("Listed Pods on all Namespace.")
}
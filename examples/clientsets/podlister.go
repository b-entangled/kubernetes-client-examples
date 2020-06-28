/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
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
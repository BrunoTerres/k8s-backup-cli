package main

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// "k8s-backup-cli/cmd"
)

// func main() {
// 	cmd.Execute()
// }

func main() {
  config, _ := clientcmd.BuildConfigFromFlags("", "/home/terres/.kube/config")
  // creates the clientset
  clientset, _ := kubernetes.NewForConfig(config)
  // access the API to list pods
  pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
  for i, s := range pods.Items.Pod {
	fmt.Println(i, s)
  }
  
}
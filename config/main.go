package config

// import (
// 	"context"


// 	"k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/tools/clientcmd"

// )

// func config() v1 {
//   config, _ := clientcmd.BuildConfigFromFlags("", "/home/terres/.kube/config")
//   // creates the clientset
//   clientset, _ := kubernetes.NewForConfig(config)
//   // access the API to list pods
//   pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
//   return pods
// }
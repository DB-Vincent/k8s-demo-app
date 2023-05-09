package main

import (
	"net/http"
	"context"
  "fmt"
//   "time"

  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/discovery"
//   "k8s.io/client-go/kubernetes"
//   "k8s.io/client-go/rest"

	"github.com/DB-Vincent/k8s-demo-app/html"
	"github.com/DB-Vincent/k8s-demo-app/utils"
)

func main() {
	http.HandleFunc("/", env)
	http.ListenAndServe(":8080", nil)
}

func env(w http.ResponseWriter, r *http.Request) {
	// Initialize environment
  	opts := &utils.KubeConfigOptions{}
    opts.Init()

  discoveryClient, err := discovery.NewDiscoveryClientForConfig(opts.Config)
  if err != nil {
      fmt.Printf(" error in discoveryClient %v",err)
  }

  serverVersion, err := discoveryClient.ServerVersion()
  if err != nil{
      fmt.Println("Error while fetching server version information", err)
  }

  pods, err := opts.Client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
  if err != nil {
    panic(err.Error())
  }

  namespaces, err := opts.Client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
  if err != nil {
    panic(err.Error())
  }

  services, err := opts.Client.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
	  panic(err.Error())
	}

	nodes, err := opts.Client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
  if err != nil {
    panic(err.Error())
  }

	p := html.EnvParams{
		Title:          "k8s-demo-app",
		Version:        fmt.Sprintf("v%s.%s (%s)", serverVersion.Major, serverVersion.Minor, serverVersion.Platform),
		PodCount:       len(pods.Items),
		NamespaceCount: len(namespaces.Items),
		ServiceCount:   len(services.Items),
		NodeCount:      len(nodes.Items),
	}
	html.Env(w, p)
}
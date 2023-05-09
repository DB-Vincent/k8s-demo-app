package utils

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
// 	"k8s.io/client-go/tools/clientcmd"
)

type KubeConfigOptions struct {
	Config *rest.Config
	Client *kubernetes.Clientset
}

func (opts *KubeConfigOptions) Init() error {
	var err error

	opts.Config, err = rest.InClusterConfig()
	if err != nil {
	  panic(err.Error())
	}

	opts.Client, err = kubernetes.NewForConfig(opts.Config)
	if err != nil {
		return err
	}

	return nil
}
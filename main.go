package main

import (
	"flag"
	"fmt"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubepath := flag.String("kubeconfig", "/home/pthangad/.kube/config", "Kubeconfig file path")
	flag.Parse()

	k, _, err := NewKubeClient(*kubepath, "")
	if err != nil {
		fmt.Errorf("failed to create kubeclient from config file at %s: %s", *kubepath, err)
	}
	ch := make(chan struct{})
	informers := informers.NewSharedInformerFactory(k, 10*time.Minute)
	c := newController(k, informers.Apps().V1().Deployments())
	informers.Start(ch)
	c.run(ch)
}

func NewKubeClient(configPath string, clusterName string) (*kubernetes.Clientset, *rest.Config, error) {
	cfg, err := BuildClientConfig(configPath, clusterName)
	if err != nil {
		cfg, err = rest.InClusterConfig()
		if err != nil {
			return nil, nil, err
		}
	}

	k, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, nil, err
	}
	return k, cfg, nil
}

// BuildClientConfig builds the client config specified by the config path and the cluster name
func BuildClientConfig(kubeConfigPath string, clusterName string) (*rest.Config, error) {
	overrides := clientcmd.ConfigOverrides{}
	// Override the cluster name if provided.
	if clusterName != "" {
		overrides.Context.Cluster = clusterName
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath},
		&overrides).ClientConfig()
}

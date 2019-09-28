package kcfcf

// Kubernetes Config from CLI Flags
//
// It is a common task to set up a client-go CLI program to read flags
// and generate *rest.Config, and this helps simplify the task.

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// KCFCF struct
type KCFCF struct {
	kubeconfig   *string
	context      *string
	inCluster    *bool
	kubeAPIQps   *float64
	kubeAPIBurst *float64
}

// NewKCFCF returns a new KCFCF structure
func NewKCFCF() *KCFCF {
	kcfcf := KCFCF{}
	return &kcfcf
}

// Init initializes flags setup
func (f *KCFCF) Init() {
	if home := homedir.HomeDir(); home != "" {
		f.kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		f.kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	f.context = flag.String("context", "", "(optional) kubeconfig file context to use")
	f.inCluster = flag.Bool("in-cluster", false, "use in-cluster config for Kubernetes API")
	f.kubeAPIQps = flag.Float64("kube-api-qps", 5.0, `QPS limit when making requests to Kubernetes apiserver`)
	f.kubeAPIBurst = flag.Float64("kube-api-burst", 10.0, `QPS burst limit when making requests to Kubernetes apiserver`)
}

// GetConfig returns parsed config from CLI flags
func (f *KCFCF) GetConfig() *rest.Config {
	var config *rest.Config
	if *f.inCluster {
		var err error
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	} else {
		var err error
		loadingRules := clientcmd.ClientConfigLoadingRules{ExplicitPath: *f.kubeconfig}
		configOverrides := clientcmd.ConfigOverrides{CurrentContext: *f.context}
		kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(&loadingRules, &configOverrides)
		config, err = kubeConfig.ClientConfig()
		if err != nil {
			panic(err.Error())
		}
	}

	config.QPS = float32(*f.kubeAPIQps)
	config.Burst = int(*f.kubeAPIBurst)

	return config
}

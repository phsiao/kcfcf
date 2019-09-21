package main

// count number of pod objets in the cluster
// using kcfcf so you can point to different clusters using --context

import (
	"flag"
	"fmt"

	kcfcf "github.com/phsiao/kcfcf"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	cliflags *kcfcf.KCFCF
)

func init() {
	cliflags = kcfcf.NewKCFCF()
	cliflags.Init()
}

func main() {
	flag.Parse()
	config := cliflags.GetConfig()

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}

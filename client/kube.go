package client

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/SwaDeshiTech/kubesync/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var K8sClientSetMap map[string]*kubernetes.Clientset

func GetClient(kubeConfigPath string) (*kubernetes.Clientset, error) {

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create a Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	return clientset, nil
}

func LoadKubernestesClients() {

	kubeConfigFolderPath := config.GetConfig().KubeConfigPath

	files, err := ioutil.ReadDir(kubeConfigFolderPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("----started loading k8s client----")

	for _, file := range files {
		k8sClientSet, err := GetClient(fmt.Sprintf("%s/%s", kubeConfigFolderPath, file.Name()))
		if err != nil {
			log.Printf("---error loading kubeconfig file for %s---%v", file.Name(), err)
			continue
		}
		K8sClientSetMap[file.Name()] = k8sClientSet
	}

	log.Println("----finished loading k8s client----")
}

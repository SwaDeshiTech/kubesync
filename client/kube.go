package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/SwaDeshiTech/kubesync/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var K8sClientSetMap map[string]*kubernetes.Clientset

func GetClient(kubeConfigPath string) (*kubernetes.Clientset, error) {

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func GetInClusterClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func LoadKubernestesClients() {

	K8sClientSetMap = make(map[string]*kubernetes.Clientset)

	log.Println("----started loading k8s client----")
	loadServiceAccountClient()
	loadKubeConfigClients()
	log.Println("----finished loading k8s client----")
}

func loadServiceAccountClient() {
	if !shouldUseServiceAccount() {
		return
	}

	clusterName := config.GetConfig().ServiceAccountName
	if clusterName == "" {
		clusterName = "in-cluster"
	}

	k8sClientSet, err := GetInClusterClient()
	if err != nil {
		log.Printf("---error loading service account client %s---%v", clusterName, err)
		return
	}

	K8sClientSetMap[clusterName] = k8sClientSet
}

func loadKubeConfigClients() {
	kubeConfigFolderPath := config.GetConfig().KubeConfigPath
	if kubeConfigFolderPath == "" {
		return
	}

	stat, err := os.Stat(kubeConfigFolderPath)
	if err != nil || !stat.IsDir() {
		log.Printf("----skipping kubeconfig loading, invalid folder %s----", kubeConfigFolderPath)
		return
	}

	files, err := ioutil.ReadDir(kubeConfigFolderPath)
	if err != nil {
		log.Printf("----failed to read kubeconfig folder %s: %v----", kubeConfigFolderPath, err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		k8sClientSet, err := GetClient(fmt.Sprintf("%s/%s", kubeConfigFolderPath, file.Name()))
		if err != nil {
			log.Printf("---error loading kubeconfig file for %s---%v", file.Name(), err)
			continue
		}
		K8sClientSetMap[file.Name()] = k8sClientSet
	}
}

func shouldUseServiceAccount() bool {
	if config.GetConfig().UseServiceAccount {
		return true
	}

	return os.Getenv("KUBERNETES_SERVICE_HOST") != "" && os.Getenv("KUBERNETES_SERVICE_PORT") != ""
}

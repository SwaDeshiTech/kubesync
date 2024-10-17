package kubernetes

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type NameSpaceWatcher struct {
	ClientSet        *kubernetes.Clientset
	NamespaceChannel chan string
}

func (namespaceWatcher *NameSpaceWatcher) Watch() {
	// Watch for changes in namespaces
	watcher, err := namespaceWatcher.ClientSet.CoreV1().Namespaces().Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the watch events
	for {
		event := <-watcher.ResultChan()
		switch event.Type {
		case watch.Added:
			namespaceName := event.Object.(*v1.Namespace).Name
			namespaceWatcher.NamespaceChannel <- namespaceName
		}
	}
}

func CreateNamespaceChannel() chan string {
	return make(chan string)
}

func SubscribeToNamespaceChannel(namespaceChannel chan string, syncResource SyncResource) {
	for {
		select {
		case namespace := <-namespaceChannel:
			fmt.Println("Received new namespace:", namespace)
			syncResource.DestinationNameSpace = namespace
			syncResource.SyncResources()
		}
	}
}

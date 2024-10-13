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
	ClientSet *kubernetes.Clientset
}

func (namespaceWatcher *NameSpaceWatcher) Watcher() {
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
			fmt.Printf("New namespace created: %s\n", event.Object.(*v1.Namespace).Name)
			syncResource := SyncResource{
				DestinationNameSpace: event.Object.(*v1.Namespace).Name,
				SourceNameSpace:      "kubesync",
			}
			syncResource.SyncResources()
		}
	}
}

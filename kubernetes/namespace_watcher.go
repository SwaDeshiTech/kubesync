package kubernetes

import (
	"context"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type NameSpaceWatcher struct {
	ClientSet *kubernetes.Clientset
	Broker    *Broker
}

func (namespaceWatcher *NameSpaceWatcher) Watch() {
	for {
		// Watch for changes in namespaces
		watcher, err := namespaceWatcher.ClientSet.CoreV1().Namespaces().Watch(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Printf("Error watching namespaces: %v. Retrying in 5 seconds...", err)
			// Retry after 5 seconds
			continue
		}

		log.Println("Namespace watcher started successfully")

		// Loop through the watch events
		for event := range watcher.ResultChan() {
			switch event.Type {
			case watch.Added:
				namespaceName := event.Object.(*v1.Namespace).Name
				log.Printf("Namespace added: %s", namespaceName)
				go namespaceWatcher.Broker.Publish("namespace", namespaceName)
			case watch.Error:
				log.Printf("Watch error received: %v", event.Object)
			}
		}

		log.Println("Namespace watcher channel closed, reconnecting...")
		watcher.Stop()
	}
}

func SubscribeToNamespaceChannel(broker *Broker, syncResource SyncResource) {

	subscriber := broker.AddSubscriber()
	broker.Subscribe(subscriber, "namespace")

	go subscriber.Listen(syncResource)
}

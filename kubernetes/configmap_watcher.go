package kubernetes

import (
	"context"
	"log"

	"github.com/SwaDeshiTech/kubesync/client"
	"github.com/SwaDeshiTech/kubesync/config"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapWatcher struct {
	ClientSet     *kubernetes.Clientset
	Namespace     string
	ConfigMapName string
	Broker        *Broker
}

func (configMapWatcher *ConfigMapWatcher) Watcher() {
	// Watch for changes in configmap
	watcher, err := configMapWatcher.ClientSet.CoreV1().ConfigMaps(configMapWatcher.ConfigMapName).Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the watch events
	for {
		event := <-watcher.ResultChan()
		switch event.Type {
		case watch.Added:
		case watch.Modified:
			configMapName := event.Object.(*v1.ConfigMap).Name
			go configMapWatcher.Broker.Publish("configmap", configMapName)
		}
	}
}

func SubscribeToConfigMapChange() {

	syncerConfigs := config.GetSyncerConfig().SyncerConfigs

	for _, itr := range syncerConfigs {
		for _, configMap := range itr.ConfigMap.List {
			configMapWatcher := ConfigMapWatcher{
				ClientSet:     client.K8sClientSetMap[itr.K8sClusterName],
				Namespace:     itr.ConfigMap.SourceNamespace,
				ConfigMapName: configMap,
			}
			configMapWatcher.Watcher()
		}
	}
}

func SubscribeToConfigMapChannel(broker *Broker, syncResource SyncResource) {

	subscriber := broker.AddSubscriber()
	broker.Subscribe(subscriber, "configmap")

	go subscriber.ListenConfigMap(syncResource)
}

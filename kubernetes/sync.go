package kubernetes

import (
	"log"
	"slices"

	"github.com/SwaDeshiTech/kubesync/client"
	"github.com/SwaDeshiTech/kubesync/config"
	"k8s.io/client-go/kubernetes"
)

type SyncResource struct {
	K8sClient            *kubernetes.Clientset
	SourceNameSpace      string
	DestinationNameSpace string
	SyncerConfig         config.Syncer
}

func SubscribeSyncResourcesToWatcher(namespaceChannel chan string) {
	for _, resource := range config.GetSyncerConfig().SyncerConfigs {
		go func() {
			k8sKubeClient, err := client.GetClient()
			if err != nil {
				log.Println("failed to obtain client set", err)
				return
			}
			syncResource := SyncResource{
				K8sClient:       k8sKubeClient,
				SourceNameSpace: resource.SourceNamespace,
				SyncerConfig:    resource,
			}
			SubscribeToNamespaceChannel(namespaceChannel, syncResource)
		}()
	}
}

func (syncResource *SyncResource) SyncResources() {

	if !slices.Contains(syncResource.SyncerConfig.DestinationNamespace, syncResource.DestinationNameSpace) {
		return
	}

	for _, syncerConfig := range config.GetSyncerConfig().SyncerConfigs {

		for _, configMapSyncer := range syncerConfig.ConfigMapList {
			configMapSyncer := SyncK8s{
				ClientSet:            syncResource.K8sClient,
				SourceNameSpace:      syncResource.SourceNameSpace,
				DestinationNameSpace: syncResource.DestinationNameSpace,
				ResourceName:         configMapSyncer,
			}
			configMapSyncer.SyncConfigMap()
		}

		for _, secretSyncer := range syncerConfig.SecretList {
			secretSyncer := SyncK8s{
				ClientSet:            syncResource.K8sClient,
				SourceNameSpace:      syncResource.SourceNameSpace,
				DestinationNameSpace: syncResource.DestinationNameSpace,
				ResourceName:         secretSyncer,
			}
			secretSyncer.SyncSecret()
		}

	}
}

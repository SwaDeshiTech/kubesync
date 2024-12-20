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

func SubscribeSyncResourcesToWatcher(broker *Broker) {
	for _, resource := range config.GetSyncerConfig().SyncerConfigs {
		go func() {
			syncResource := SyncResource{
				K8sClient:       client.K8sClientSetMap[resource.K8sClusterName],
				SourceNameSpace: resource.SourceNamespace,
				SyncerConfig:    resource,
			}
			SubscribeToNamespaceChannel(broker, syncResource)
		}()
	}
}

func (syncResource *SyncResource) SyncResources() {

	if !slices.Contains(syncResource.SyncerConfig.DestinationNamespace, syncResource.DestinationNameSpace) {
		return
	}

	log.Printf("----Executing syncer %s syncing resource from namespace %s to %s----", syncResource.SyncerConfig.Name, syncResource.SourceNameSpace, syncResource.DestinationNameSpace)

	for _, configMapSyncer := range syncResource.SyncerConfig.ConfigMapList {
		configMapSyncer := SyncK8s{
			ClientSet:            syncResource.K8sClient,
			SourceNameSpace:      syncResource.SourceNameSpace,
			DestinationNameSpace: syncResource.DestinationNameSpace,
			ResourceName:         configMapSyncer,
		}
		configMapSyncer.SyncConfigMap()
	}

	for _, secretSyncer := range syncResource.SyncerConfig.SecretList {
		secretSyncer := SyncK8s{
			ClientSet:            syncResource.K8sClient,
			SourceNameSpace:      syncResource.SourceNameSpace,
			DestinationNameSpace: syncResource.DestinationNameSpace,
			ResourceName:         secretSyncer,
		}
		secretSyncer.SyncSecret()
	}

}

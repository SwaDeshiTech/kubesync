package kubernetes

import (
	"log"

	"github.com/SwaDeshiTech/kubesync/client"
	"github.com/SwaDeshiTech/kubesync/config"
)

type SyncResource struct {
	SourceNameSpace      string
	DestinationNameSpace string
}

func (syncResource *SyncResource) SyncResources() {

	k8sKubeClient, err := client.GetClient()
	if err != nil {
		log.Println("failed to obtain client set", err)
		return
	}

	for _, configMapSyncer := range config.GetSyncerConfig().ConfigMapList {
		configMapSyncer := SyncK8s{
			ClientSet:            k8sKubeClient,
			SourceNameSpace:      syncResource.SourceNameSpace,
			DestinationNameSpace: syncResource.DestinationNameSpace,
			ResourceName:         configMapSyncer,
		}
		configMapSyncer.SyncConfigMap()
	}

	for _, secretSyncer := range config.GetSyncerConfig().SecretList {
		secretSyncer := SyncK8s{
			ClientSet:            k8sKubeClient,
			SourceNameSpace:      syncResource.SourceNameSpace,
			DestinationNameSpace: syncResource.DestinationNameSpace,
			ResourceName:         secretSyncer,
		}
		secretSyncer.SyncSecret()
	}
}

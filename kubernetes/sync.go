package kubernetes

import (
	"log"

	"github.com/SwaDeshiTech/kubesync/client"
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

	configMapSyncer := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "outbound",
	}

	configMapSyncer.SyncConfigMap()

	secretSyncer := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "jfrog",
	}

	secretSyncer.SyncSecret()
}

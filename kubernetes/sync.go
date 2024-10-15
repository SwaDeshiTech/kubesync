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

	configMapSyncer1 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "outboundaddress",
	}
	configMapSyncer1.SyncConfigMap()

	configMapSyncer2 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "confs",
	}
	configMapSyncer2.SyncConfigMap()

	configMapSyncer3 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "coredns",
	}
	configMapSyncer3.SyncConfigMap()

	configMapSyncer4 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "corednstmpl",
	}
	configMapSyncer4.SyncConfigMap()

	configMapSyncer5 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "blockedservices",
	}
	configMapSyncer5.SyncConfigMap()

	configMapSyncer6 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "dns-config",
	}
	configMapSyncer6.SyncConfigMap()

	configMapSyncer7 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "kube-root-ca.crt",
	}
	configMapSyncer7.SyncConfigMap()

	secretSyncer1 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "jfrog",
	}
	secretSyncer1.SyncSecret()

	secretSyncer2 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "az-sedockerreg",
	}
	secretSyncer2.SyncSecret()

	secretSyncer3 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "az-sejfrogreg",
	}
	secretSyncer3.SyncSecret()

	secretSyncer4 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "az-sejfrogregci",
	}
	secretSyncer4.SyncSecret()

	secretSyncer5 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "certs",
	}
	secretSyncer5.SyncSecret()

	secretSyncer6 := SyncK8s{
		ClientSet:            k8sKubeClient,
		SourceNameSpace:      syncResource.SourceNameSpace,
		DestinationNameSpace: syncResource.DestinationNameSpace,
		ResourceName:         "dockerjfrog-prodhub",
	}
	secretSyncer6.SyncSecret()
}

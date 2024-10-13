package kubernetes

import (
	"log"

	"k8s.io/client-go/kubernetes"
)

type SyncK8s struct {
	ClientSet            *kubernetes.Clientset
	SourceNameSpace      string
	DestinationNameSpace string
	ResourceName         string
}

func (k8sSyncer *SyncK8s) SyncConfigMap() {
	err := SyncConfigMap(k8sSyncer.ClientSet, k8sSyncer.SourceNameSpace, k8sSyncer.DestinationNameSpace, k8sSyncer.ResourceName)
	if err != nil {
		log.Println("----error in syncing config map----", err)
	}
}

func (k8sSyncer *SyncK8s) SyncSecret() {
	err := SyncSecret(k8sSyncer.ClientSet, k8sSyncer.SourceNameSpace, k8sSyncer.DestinationNameSpace, k8sSyncer.ResourceName)
	if err != nil {
		log.Println("----error in syncing secret----", err)
	}
}

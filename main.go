package main

import (
	"fmt"
	"log"

	"github.com/SwaDeshiTech/kubesync/api"
	"github.com/SwaDeshiTech/kubesync/client"
	"github.com/SwaDeshiTech/kubesync/config"
	"github.com/SwaDeshiTech/kubesync/cron"
	"github.com/SwaDeshiTech/kubesync/databases"
	"github.com/SwaDeshiTech/kubesync/kubernetes"
)

func main() {

	//load config.yml
	if err := config.ReadConfig(); err != nil {
		panic(err)
	}

	//load sync config.yml
	if err := config.ReadSyncerConfig(); err != nil {
		panic(err)
	}

	if !config.GetConfig().DisableCronJob {
		if err := databases.InitializeMongoConnection(); err != nil {
			panic(err)
		}

		go func() {
			cron.InitializeCrons()
		}()
	}

	namespaceChannel := kubernetes.CreateNamespaceChannel()

	go func() {
		log.Println("----Starting namespace watcher----")
		k8sKubeClient, err := client.GetClient()
		if err != nil {
			log.Println("failed to obtain client set", err)
			return
		}

		namespaceWatcher := kubernetes.NameSpaceWatcher{
			ClientSet:        k8sKubeClient,
			NamespaceChannel: namespaceChannel,
		}

		namespaceWatcher.Watch()
	}()

	go func() {
		log.Println("----Subscribing sync resources to watcher----")
		kubernetes.SubscribeSyncResourcesToWatcher(namespaceChannel)
		log.Println("----Completed Subscribing sync resources to watcher----")
	}()

	router := api.ServerV1()
	if err := router.Run(fmt.Sprintf(":%d", config.GetConfig().Port)); err != nil {
		panic(err)
	}
}

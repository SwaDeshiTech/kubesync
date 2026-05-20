package main

import (
	"log"

	"github.com/SwaDeshiTech/kubesync/client"
	"github.com/SwaDeshiTech/kubesync/config"
	"github.com/SwaDeshiTech/kubesync/cron"
	"github.com/SwaDeshiTech/kubesync/kubernetes"
)

func main() {

	//load config.yml
	if err := config.ReadConfig(); err != nil {
		panic(err)
	}

	client.LoadKubernestesClients()

	if !config.GetConfig().DisableCronJob {
		go func() {
			cron.InitializeCrons()
		}()
	}

	broker := kubernetes.NewBroker()

	go func() {
		log.Println("----Starting namespace watcher----")
		for _, k8sClientSet := range client.K8sClientSetMap {
			namespaceWatcher := kubernetes.NameSpaceWatcher{
				ClientSet: k8sClientSet,
				Broker:    broker,
			}
			namespaceWatcher.Watch()
		}
	}()

	go func() {
		log.Println("----Subscribing sync resources to watcher----")
		kubernetes.SubscribeSyncResourcesToWatcher(broker)
		log.Println("----Completed Subscribing sync resources to watcher----")
	}()

	select {}
}

package cron

import (
	"context"
	"log"

	"github.com/SwaDeshiTech/kubesync/enums"
	"github.com/SwaDeshiTech/kubesync/services/job"
)

func InitializeCrons() error {

	ctx := context.Background()

	handlers := initializeCronHandlers()

	factory := job.NewCronFactory(handlers)

	for _, priority := range enums.P0.Values() {
		cronGroup, err := factory.NewCronGroup(priority)
		if err != nil {
			log.Println("failed to get cron group from factory", err)
		}

		log.Printf("starting cron group %s", cronGroup.CronGroupName)

		cronGroup.InitializeCrons(ctx)
	}

	return nil
}

func initializeCronHandlers() map[string]job.CronHandler {
	return map[string]job.CronHandler{
		"SyncKubernetesResourcesJob": job.SyncKubernetesResourcesJob{},
	}
}

package job

import (
	"log"

	v1 "github.com/SwaDeshiTech/arsenal/pkg/mongo-connector/v1"
	"github.com/SwaDeshiTech/kubesync/entity"
	"github.com/SwaDeshiTech/kubesync/enums"
	"github.com/SwaDeshiTech/kubesync/services"
	"go.mongodb.org/mongo-driver/bson"
)

type SyncKubernetesResourcesJob struct {
	JobId   string `json:"jobId"`
	JobName string `json:"jobName"`
}

func (j SyncKubernetesResourcesJob) ExecuteCron(cronId string) error {

	filters := bson.M{"uuid": cronId}
	sort := bson.D{{Key: "name", Value: 1}}

	resultCriteria := v1.ResultCriteria{
		Filters: filters,
		Sort:    sort,
	}

	jobDetailConfig, err := entity.FetchCronScheduleConfigs(resultCriteria)
	if err != nil || len(jobDetailConfig) == 0 {
		log.Println("cron job config could not fetched", err)
		return err
	}

	services.UpdateCronSchedule(jobDetailConfig[0].UUID, string(enums.Running))

	return nil
}

func (j SyncKubernetesResourcesJob) StopCron(eventId string) error {
	return nil
}

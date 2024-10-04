package services

import (
	"log"
	"net/http"
	"time"

	v1 "github.com/SwaDeshiTech/arsenal/pkg/mongo-connector/v1"
	customerror "github.com/SwaDeshiTech/kubesync/customError"
	"github.com/SwaDeshiTech/kubesync/dto"
	"github.com/SwaDeshiTech/kubesync/entity"
	"github.com/SwaDeshiTech/kubesync/enums"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddCronSchedule(cronScheduleConfigRequest dto.CronScheduleConfigRequest) (dto.CronScheduleConfigResponse, customerror.Http) {

	var cronSchedule entity.CronSchedule
	err := mapstructure.Decode(cronScheduleConfigRequest, &cronSchedule)
	if err != nil {
		log.Println("failed to map the dto & entity", err)
		return dto.CronScheduleConfigResponse{}, customerror.HttpError("Failed to map the DTO & Entity", http.StatusBadRequest)
	}

	cronSchedule.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	cronSchedule.Status = enums.Pending

	err = cronSchedule.Insert()
	if err != nil {
		log.Println("failed to insert k8sCluster config", err)
		return dto.CronScheduleConfigResponse{}, customerror.HttpError("Failed to onbaord Cron Schedule", http.StatusInternalServerError)
	}

	var cronScheduleConfigResponse dto.CronScheduleConfigResponse
	err = mapstructure.Decode(cronSchedule, &cronScheduleConfigResponse)
	if err != nil {
		log.Println("Failed to map Entity & DTO", err)
		return dto.CronScheduleConfigResponse{}, customerror.HttpError("Failed to map Entity & DTO", http.StatusInternalServerError)
	}

	return cronScheduleConfigResponse, customerror.Http{}
}

func GetCronSchedule(cronScheduleName string) (dto.CronScheduleConfigResponse, customerror.Http) {

	filters := bson.M{"name": cronScheduleName}
	sort := bson.D{{Key: "name", Value: 1}}

	resultCriteria := v1.ResultCriteria{
		Filters: filters,
		Sort:    sort,
	}

	var cronScheduleConfigResponse dto.CronScheduleConfigResponse

	cronScheduleConfigs, err := entity.FetchCronScheduleConfigs(resultCriteria)
	if err != nil {
		log.Println("failed to fetch cron schedule config", err)
		return cronScheduleConfigResponse, customerror.HttpError("cron schedule config could not be fetched", http.StatusInternalServerError)
	}

	if len(cronScheduleConfigs) == 0 {
		log.Println("cron schedule config could not be found")
		return cronScheduleConfigResponse, customerror.HttpError("cron schedule config could not be found", http.StatusNotFound)
	}

	err = mapstructure.Decode(cronScheduleConfigs[0], &cronScheduleConfigResponse)
	if err != nil {
		log.Println("failed to map entity & dto", err)
		return cronScheduleConfigResponse, customerror.HttpError("cron schedule config could not be fetched", http.StatusInternalServerError)
	}

	return cronScheduleConfigResponse, customerror.Http{}
}

func UpdateCronSchedule(id string, status string) customerror.Http {

	resultCriteria := v1.ResultCriteria{
		Filters: primitive.M{
			"id": id,
		},
		Update: primitive.M{"$set": bson.M{
			"status": status,
		}},
	}

	err := entity.UpdateCronSchedule(resultCriteria)
	if err != nil {
		log.Println("failed to update the status of request", id, status)
		return customerror.HttpError("Status could not be updated", http.StatusInternalServerError)
	}

	return customerror.Http{}
}

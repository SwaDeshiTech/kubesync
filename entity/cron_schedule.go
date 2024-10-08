package entity

import (
	"context"

	jsonV1 "github.com/SwaDeshiTech/arsenal/pkg/json/v1"
	v1 "github.com/SwaDeshiTech/arsenal/pkg/mongo-connector/v1"
	"github.com/SwaDeshiTech/kubesync/constants"
	"github.com/SwaDeshiTech/kubesync/databases"
	"github.com/SwaDeshiTech/kubesync/enums"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CronSchedule struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UUID           string             `json:"uuid" bson:"uuid"`
	CronExpression string             `json:"cronExpression" bson:"cronExpression"`
	Status         enums.Status       `json:"status" bson:"status"`
	JobName        string             `json:"jobName" bson:"jobName"`
	JobDescription string             `json:"jobDescription" bson:"jobDescription"`
	StartDate      string             `json:"startDate" bson:"startDate"`
	EndDate        string             `json:"endDate" bson:"endDate"`
	Priority       enums.Priority     `json:"priority" bson:"priority"`
	Frequency      enums.Frequency    `json:"frequency" bson:"frequency"`
	IsActive       bool               `json:"isActive" bson:"isActive"`
	JobType        string             `json:"jobType" bson:"jobType"`
	CreatedAt      primitive.DateTime `json:"createdAt" bson:"createdAt"`
	Resources      []Resource         `json:"resources" bson:"resources"`
}

type Resource struct {
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
}

func (cronSchedule *CronSchedule) Insert() error {

	mongoClient := databases.MongoClient

	db := mongoClient.Database(constants.MONGO_KUBE_SYNC_DB_NAME)
	cronScheduleCollection := db.Collection(constants.CRON_SCHEDULE_COLLECTION_NAME)

	_, err := cronScheduleCollection.InsertOne(context.TODO(), cronSchedule)
	if err != nil {
		return err
	}
	return nil
}

func FetchCronScheduleConfigs(resultCriteria v1.ResultCriteria) ([]CronSchedule, error) {

	mongoClient := databases.MongoClient

	db := mongoClient.Database(constants.MONGO_KUBE_SYNC_DB_NAME)
	cronScheduleCollection := db.Collection(constants.CRON_SCHEDULE_COLLECTION_NAME)

	if resultCriteria.Sort == nil {
		resultCriteria.Sort = primitive.D{}
	}

	findOptions := options.Find()
	findOptions.SetSort(resultCriteria.Sort)

	cursor, err := cronScheduleCollection.Find(context.TODO(), resultCriteria.Filters, findOptions)
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	jsonResp, err := jsonV1.ConvertIntoJSON(results)
	if err != nil {
		return nil, err
	}

	cronScheduleConfigList := []CronSchedule{}

	err = jsonV1.ParseJSON(jsonResp, &cronScheduleConfigList)
	if err != nil {
		return nil, err
	}

	return cronScheduleConfigList, nil
}

func UpdateCronSchedule(resultCriteria v1.ResultCriteria) error {

	mongoClient := databases.MongoClient

	db := mongoClient.Database(constants.MONGO_KUBE_SYNC_DB_NAME)
	cronScheduleCollection := db.Collection(constants.CRON_SCHEDULE_COLLECTION_NAME)

	_, err := cronScheduleCollection.UpdateOne(context.TODO(), resultCriteria.Filters, resultCriteria.Update)
	if err != nil {
		return err
	}
	return nil
}

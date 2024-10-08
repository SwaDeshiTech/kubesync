package dto

import (
	"github.com/SwaDeshiTech/kubesync/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CronScheduleConfigResponse struct {
	UUID           string             `json:"uuid"`
	CronExpression string             `json:"cronExpression"`
	Status         enums.Status       `json:"status"`
	JobName        string             `json:"jobName"`
	JobDescription string             `json:"jobDescription"`
	StartDate      string             `json:"startDate"`
	EndDate        string             `json:"endDate"`
	Frequency      enums.Frequency    `json:"frequency"`
	Priority       enums.Priority     `json:"priority"`
	CreatedAt      primitive.DateTime `json:"createdAt"`
	JobType        string             `json:"jobType"`
	Resources      []ResourceResponse `json:"resources"`
}

type ResourceResponse struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

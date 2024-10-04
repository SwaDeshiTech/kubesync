package dto

import (
	"github.com/SwaDeshiTech/kubesync/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CronScheduleConfigResponse struct {
	CronExpression      string             `json:"cronExpression"`
	Status              enums.Status       `json:"status"`
	JobName             string             `json:"jobName"`
	JobDescription      string             `json:"jobDescription"`
	KubernetesResources string             `json:"kubernetesResources"`
	CreatedAt           primitive.DateTime `json:"createdAt"`
}

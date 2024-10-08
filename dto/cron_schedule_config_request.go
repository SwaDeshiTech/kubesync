package dto

import "github.com/SwaDeshiTech/kubesync/enums"

type CronScheduleConfigRequest struct {
	CronExpression string          `json:"cronExpression"`
	JobName        string          `json:"jobName"`
	JobDescription string          `json:"jobDescription"`
	StartDate      string          `json:"startDate"`
	EndDate        string          `json:"endDate"`
	Frequency      enums.Frequency `json:"frequency"`
	Priority       enums.Priority  `json:"priority"`
	JobType        string          `json:"jobType"`
	Resources      []Resource      `json:"resources"`
}

type Resource struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CronScheduleConfigParam struct {
	ID string `uri:"id,required"`
}

package dto

type CronScheduleConfigRequest struct {
	CronExpression string `json:"cronExpression"`
	JobName        string `json:"jobName"`
	JobDescription string `json:"jobDescription"`
}

type CronScheduleConfigParam struct {
	ID string `uri:"id,required"`
}

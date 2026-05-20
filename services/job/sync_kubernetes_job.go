package job

import "log"

type SyncKubernetesResourcesJob struct {
	JobId   string `json:"jobId"`
	JobName string `json:"jobName"`
}

func (j SyncKubernetesResourcesJob) ExecuteCron(cronId string) error {
	log.Printf("Executing cron %s", cronId)
	return nil
}

func (j SyncKubernetesResourcesJob) StopCron(eventId string) error {
	return nil
}

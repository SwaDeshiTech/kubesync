package job

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/SwaDeshiTech/kubesync/config"
	"github.com/robfig/cron"
)

type CronJob interface {
	InitializeCrons(ctx context.Context) error
	Stop() error
}

type CronHandler interface {
	ExecuteCron(cronId string) error
	StopCron(eventId string) error
}

type CronFactory struct {
	Handlers map[string]CronHandler
}

type CronConsumer struct {
	CronGroupName string
	CronGroup     cron.Cron
	Handlers      map[string]CronHandler
}

func NewCronFactory(handlers map[string]CronHandler) *CronFactory {
	return &CronFactory{
		Handlers: handlers,
	}
}

func (f *CronFactory) NewCronGroup(priority string) (CronConsumer, error) {

	newCronGroup := cron.New()

	return CronConsumer{
		CronGroupName: priority,
		CronGroup:     *newCronGroup,
		Handlers:      f.Handlers,
	}, nil
}

func (c *CronConsumer) InitializeCrons(ctx context.Context) error {

	jobList := config.GetConfig().CronSchedules
	filtered := make([]config.CronSchedule, 0, len(jobList))
	for _, job := range jobList {
		if !job.IsActive || job.Priority != c.CronGroupName {
			continue
		}
		filtered = append(filtered, job)
	}

	if len(filtered) == 0 {
		log.Println("no job to schedule")
		return errors.New(fmt.Sprintf("job list is empty for %s", c.CronGroupName))
	}

	log.Println("-----started scheduling cron jobs-----")

	for _, job := range filtered {
		c.CronGroup.AddFunc(job.CronExpression, func() {
			err := c.Handlers[job.JobType].ExecuteCron(job.UUID)
			if err != nil {
				log.Println("failed to execute the cron", err)
			}
		})
	}

	c.CronGroup.Start()
	log.Println("-----finish scheduling cron jobs-----")

	return nil
}

func (c *CronConsumer) Stop() error {
	// Close all crons
	return nil
}

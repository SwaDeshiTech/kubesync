package job

import (
	"context"
	"errors"
	"fmt"
	"log"

	v1 "github.com/SwaDeshiTech/arsenal/pkg/mongo-connector/v1"
	"github.com/SwaDeshiTech/kubesync/dto"
	"github.com/SwaDeshiTech/kubesync/entity"
	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
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

func (f *CronFactory) NewCronGroup(config dto.CronConfig) (CronConsumer, error) {

	newCronGroup := cron.New()

	return CronConsumer{
		CronGroupName: config.CronGroupName,
		CronGroup:     *newCronGroup,
		Handlers:      f.Handlers,
	}, nil
}

func (c *CronConsumer) InitializeCrons(ctx context.Context) error {

	filters := bson.M{"isActive": true, "priority": c.CronGroupName}
	sort := bson.D{{Key: "priority", Value: 1}}

	resultCriteria := v1.ResultCriteria{
		Filters: filters,
		Sort:    sort,
	}

	jobList, err := entity.FetchCronScheduleConfigs(resultCriteria)
	if err != nil {
		log.Println("failed to fetch the list of cron job", err)
		return err
	}

	if len(jobList) == 0 {
		log.Println("no job to schedule")
		return errors.New(fmt.Sprintf("job list is empty for %s", c.CronGroupName))
	}

	log.Println("-----started scheduling cron jobs-----")

	for _, job := range jobList {
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

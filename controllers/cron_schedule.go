package controllers

import (
	"log"
	"net/http"

	customerror "github.com/SwaDeshiTech/kubesync/customError"
	"github.com/SwaDeshiTech/kubesync/dto"
	"github.com/SwaDeshiTech/kubesync/services"
	"github.com/gin-gonic/gin"
)

func AddCronSchedule(c *gin.Context) {

	var cronScheduleRequest dto.CronScheduleConfigRequest
	err := c.ShouldBindJSON(&cronScheduleRequest)
	if err != nil {
		log.Println("failed to parse payload", err)
		c.Error(customerror.HttpError("Failed to parse payload", http.StatusBadRequest))
		return
	}

	cronSchedule, cronScheduleErr := services.AddCronSchedule(cronScheduleRequest)
	if cronScheduleErr.Description != "" {
		c.Error(cronScheduleErr)
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		HttpStatus: http.StatusCreated,
		Message:    "Cron schedule has been onboarded successfully",
		Response:   cronSchedule,
	})
}

func GetCronSchedule(c *gin.Context) {

	var cronScheduleConfigParam dto.CronScheduleConfigParam
	err := c.ShouldBindUri(&cronScheduleConfigParam)
	if err != nil {
		log.Println("failed to parse the param", err)
		c.Error(customerror.HttpError("Failed to parse the param", http.StatusBadRequest))
		return
	}

	cronScheduleConfig, cronScheduleConfigErr := services.GetCronSchedule(cronScheduleConfigParam.ID)
	if cronScheduleConfigErr.StatusCode != 0 && cronScheduleConfigErr.StatusCode != 200 {
		c.Error(cronScheduleConfigErr)
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		HttpStatus: http.StatusOK,
		Message:    "Successfully fetched detail of Cron Schedule Config",
		Response:   cronScheduleConfig,
	})
}

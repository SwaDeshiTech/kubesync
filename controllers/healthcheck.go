package controllers

import (
	"net/http"

	"github.com/SwaDeshiTech/kubesync/dto"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Response{
		HttpStatus: http.StatusOK,
		Message:    "OK",
	})
}

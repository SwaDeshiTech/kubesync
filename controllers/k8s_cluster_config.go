package controllers

import (
	"log"
	"net/http"

	customerror "github.com/SwaDeshiTech/kubesync/customError"
	"github.com/SwaDeshiTech/kubesync/dto"
	"github.com/SwaDeshiTech/kubesync/services"
	"github.com/gin-gonic/gin"
)

func AddK8sClusterConfig(c *gin.Context) {

	var k8sClusterConfig dto.K8sClusterConfigRequest
	err := c.ShouldBindJSON(&k8sClusterConfig)
	if err != nil {
		log.Println("failed to parse payload", err)
		c.Error(customerror.HttpError("Failed to parse payload", http.StatusBadRequest))
		return
	}

	k8sCluster, k8sClusterOnboardErr := services.AddK8sClusterConfig(k8sClusterConfig)
	if k8sClusterOnboardErr.Description != "" {
		c.Error(k8sClusterOnboardErr)
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		HttpStatus: http.StatusCreated,
		Message:    "Successfully onboarded k8s cluster",
		Response:   k8sCluster,
	})
}

func GetK8sClusterConfig(c *gin.Context) {

	var k8sClusterName dto.K8sClusterConfigParam
	err := c.ShouldBindUri(&k8sClusterName)
	if err != nil {
		log.Println("failed to parse the param", err)
		c.Error(customerror.HttpError("Failed to parse the param", http.StatusBadRequest))
		return
	}

	k8sClusterConfig, k8sClusterDetailErr := services.GetK8sClusterConfigDetail(k8sClusterName.Name)
	if k8sClusterDetailErr.StatusCode != 0 && k8sClusterDetailErr.StatusCode != 200 {
		c.Error(k8sClusterDetailErr)
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		HttpStatus: http.StatusOK,
		Message:    "Successfully fetched detail of K8s Cluster Config",
		Response:   k8sClusterConfig,
	})
}

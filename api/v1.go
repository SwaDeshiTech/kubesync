package api

import (
	"github.com/SwaDeshiTech/kubesync/controllers"
	"github.com/gin-gonic/gin"
)

func ServerV1() *gin.Engine {

	router := gin.Default()

	RESTAPIV1 := router.Group("/api/v1")
	{
		RESTAPIV1.GET("/status", controllers.HealthCheck)
		RESTAPIV1.POST("/k8sClusterConfig", controllers.AddK8sClusterConfig)
		RESTAPIV1.GET("/k8sClusterConfig/:name", controllers.GetK8sClusterConfig)
		RESTAPIV1.POST("/cronSchedule", controllers.AddCronSchedule)
		RESTAPIV1.GET("/cronSchedule/:id", controllers.GetCronSchedule)
	}

	return router
}

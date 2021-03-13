package router

import (
	"github.com/gin-gonic/gin"
	"healthy-api/controllers"
)

func WeightRouter(group *gin.RouterGroup) {

	controller := controllers.NewWeightController()

	group.GET("/", controller.Index)
	group.GET("/:id", controller.Show)
	group.PUT("/:id", controller.Update)
	group.POST("/", controller.Store)
	group.DELETE("/:id", controller.Delete)
}
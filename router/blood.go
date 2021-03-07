package router

import (
	"github.com/gin-gonic/gin"
	"healthy-api/controllers"
)

func BloodRouter(group *gin.RouterGroup) {

	controller := controllers.NewBloodController()

	group.GET("/", controller.Index)
	group.DELETE("/:id", controller.Delete)
	group.GET("/:id", controller.Show)
	group.PUT("/:id", controller.Update)
	group.POST("/", controller.Store)
}
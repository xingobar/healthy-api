package router

import (
	"github.com/gin-gonic/gin"
	"healthy-api/controllers"
)

func BloodRouter(group *gin.RouterGroup) {

	controller := controllers.NewBloodController()

	group.GET("/", controller.Index)
}
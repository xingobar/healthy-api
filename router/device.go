package router

import (
	"github.com/gin-gonic/gin"
	"healthy-api/controllers"
)

func DeviceRouter(group *gin.RouterGroup) {

	bloodController := controllers.NewBloodController()

	deviceGroup := group.Group("/:deviceid")

	deviceGroup.GET("/blood", bloodController.Index)
}

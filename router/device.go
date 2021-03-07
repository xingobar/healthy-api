package router

import (
	"github.com/gin-gonic/gin"
	"healthy-api/controllers"
)

func DeviceRouter(group *gin.RouterGroup) {

	bloodController := controllers.NewBloodController()

	deviceGroup := group.Group("/:deviceid")

	deviceGroup.GET("/bloods", bloodController.Index)
	deviceGroup.DELETE("/bloods/:id", bloodController.Delete)
}

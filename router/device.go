package router

import (
	"github.com/gin-gonic/gin"
)

func DeviceRouter(group *gin.RouterGroup) {

	deviceGroup := group.Group("/:deviceid")

	bloodGroup := deviceGroup.Group("/bloods")
	BloodRouter(bloodGroup)
}

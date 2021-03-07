package router

import (
	"github.com/gin-gonic/gin"
)

func DeviceRouter(group *gin.RouterGroup) {

	deviceGroup := group.Group("/:deviceid")

	// 血壓
	bloodGroup := deviceGroup.Group("/bloods")
	BloodRouter(bloodGroup)

	// 體重
	weightGroup := deviceGroup.Group("/weights")
	WeightRouter(weightGroup)
}

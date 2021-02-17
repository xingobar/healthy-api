package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BloodRouter(group *gin.RouterGroup) {
	group.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
		})
	})
}
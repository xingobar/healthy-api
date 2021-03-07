package controllers

import (
	"github.com/gin-gonic/gin"
	"healthy-api/models"
	"net/http"
)

type bloodController struct {

}

func NewController() *bloodController{
	return &bloodController{}
}

func (c *bloodController) Index(ctx *gin.Context) {
	var bloods []models.Blood
	models.Db.Model(&models.Blood{}).Find(&bloods)

	ctx.JSON(http.StatusOK, bloods)
}

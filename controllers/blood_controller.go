package controllers

import (
	"github.com/gin-gonic/gin"
	"healthy-api/models"
	"net/http"
	"strconv"
)

type bloodController struct {

}

func NewBloodController() *bloodController{
	return &bloodController{}
}

func (c *bloodController) Index(ctx *gin.Context) {
	var bloods []models.Blood
	models.Db.Model(&models.Blood{}).
		Scopes(models.Paginate(ctx.Request)).
		Where("device_id = ?", ctx.Param("deviceid")).
		Find(&bloods)

	total := 0

	models.Db.Model(&models.Blood{}).Where("device_id = ?", ctx.Param("deviceid")).Count(&total)

	page, err := strconv.Atoi(ctx.Request.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bloods,
		"meta": map[string]interface{} {
			"total": total,
			"page": page,
		},
	})
}

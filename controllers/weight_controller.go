package controllers

import (
	"github.com/gin-gonic/gin"
	"healthy-api/models"
	"net/http"
	"strconv"
)

type weightController struct {

}

func NewWeightController() *weightController {
	return &weightController{}
}

// 首頁資料
func (c *weightController) Index(ctx *gin.Context) {

	var weights []models.Weight

	if err := models.Db.Model(&models.Weight{}).
		Scopes(models.Paginate(ctx.Request)).
		Find(&weights).Error; err != nil {

			ctx.JSON(http.StatusBadGateway, gin.H{})
			return
	}

	total := 0
	if err := models.Db.Model(&models.Weight{}).
		Where("device_id = ?", ctx.Param("deviceid")).
		Count(&total).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{})
	}

	page, err := strconv.Atoi(ctx.Request.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": weights,
		"meta": map[string]interface{}{
			"total": total,
			"page": page,
		},
	})
}
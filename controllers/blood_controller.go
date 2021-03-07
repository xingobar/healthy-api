package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"healthy-api/models"
	"net/http"
	"strconv"
)

type bloodController struct {

}

func NewBloodController() *bloodController{
	return &bloodController{}
}

// 血壓列表
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

// 刪除
func (c *bloodController) Delete(ctx *gin.Context) {
	var blood models.Blood

	if err := models.Db.Model(&models.Blood{}).
			Scopes((&models.Blood{}).GetBlood(ctx.Param("deviceid"), ctx.Param("id"))).
			First(&blood).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gorm.ErrRecordNotFound)
	} else {
		models.Db.Delete(&models.Blood{},ctx.Param("id"))
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

// 顯示血壓紀錄
func (c *bloodController) Show(ctx *gin.Context) {
	var blood models.Blood

	if err := models.Db.Model(&models.Blood{}).
		Scopes((&models.Blood{}).GetBlood(ctx.Param("deviceid"), ctx.Param("id"))).
		First(&blood).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gorm.ErrRecordNotFound)
		return
	}

	ctx.JSON(http.StatusOK, blood)
}
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
	"healthy-api/models"
	"healthy-api/validation"
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

func (c *bloodController) Update(ctx *gin.Context) {

	var blood models.Blood

	var json struct{
		Pulse int  `json:"pulse" form:"id" binding:"required"`
		Diastolic float32 `json:"diastolic" form:"diastolic" binding:"required"`
		Systolic float32 `json:"systolic" form:"systolic" binding:"required"`
	}

	message := map[string]string{
		"Pulse.required": "請輸入脈搏資料",
		"Diastolic.required": "請輸入舒張壓",
		"Systolic.required": "請輸入收縮壓",
	}

	if err := models.Db.Model(&models.Blood{}).
		Scopes((&models.Blood{}).GetBlood(ctx.Param("deviceid"), ctx.Param("id"))).
		First(&blood).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gorm.ErrRecordNotFound)

			return
	}

    if err := ctx.ShouldBindJSON(&json); err != nil {
    	fmt.Println(json)

    	ctx.JSON(http.StatusBadRequest, gin.H{
    		"msg": validation.GetError(err.(validator.ValidationErrors), message),
		})
    	return
	}

    if err := models.Db.Model(&blood).Update(json).Error; err != nil {
    	ctx.JSON(http.StatusBadGateway, gin.H{})
    	return
	}

    ctx.JSON(http.StatusOK, gin.H{})
}
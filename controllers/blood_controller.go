package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
	"healthy-api/models"
	"healthy-api/validation"
	"healthy-api/validation/blood_validation"
	"math"
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

	limit, _ := strconv.Atoi(ctx.Request.URL.Query().Get("limit"))

	if limit == 0 {
		limit = 20
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bloods,
		"meta": map[string]interface{} {
			"total": total,
			"page": page,
			"total_page": int(math.Ceil(float64(total) / float64(limit))),
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

// 更新
func (c *bloodController) Update(ctx *gin.Context) {

	var blood models.Blood

	v := blood_validation.NewValidation()

	if err := models.Db.Model(&models.Blood{}).
		Scopes((&models.Blood{}).GetBlood(ctx.Param("deviceid"), ctx.Param("id"))).
		First(&blood).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gorm.ErrRecordNotFound)

			return
	}

    json := blood_validation.GetUpdateRule{}

    if err := ctx.ShouldBindJSON(&json); err != nil {
    	ctx.JSON(http.StatusBadRequest, gin.H{
    		"msg": validation.GetError(err.(validator.ValidationErrors), v.GetMessage()),
		})
    	return
	}

    if err := models.Db.Model(&blood).Update(json).Error; err != nil {
    	ctx.JSON(http.StatusBadGateway, gin.H{})
    	return
	}

    ctx.JSON(http.StatusOK, blood)
}

func (c *bloodController) Store(ctx *gin.Context) {

	v := blood_validation.NewValidation()

	json := blood_validation.GetStoreRule{}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": validation.GetError(err.(validator.ValidationErrors), v.GetMessage()),
		})
		return
	}

	var blood = models.Blood{
		Pulse: json.Pulse,
		Diastolic:json.Diastolic,
		Systolic: json.Systolic,
		DeviceId: ctx.Param("deviceid"),
	}

	if err := models.Db.Create(&blood).Error; err != nil {

		ctx.JSON(http.StatusBadGateway, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, blood)
}
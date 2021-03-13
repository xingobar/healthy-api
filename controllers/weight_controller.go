package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
	"healthy-api/models"
	"healthy-api/validation"
	"healthy-api/validation/weight_validation"
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

//  顯示
func (c *weightController) Show(ctx *gin.Context) {
	var weight models.Weight

	if err := models.Db.Model(&models.Weight{}).
		Scopes((&models.Weight{}).GetWeight(ctx.Param("deviceid"), ctx.Param("id"))).
		First(&weight).Error; err != nil {
			ctx.JSON(http.StatusNotExtended, gorm.ErrRecordNotFound)
			return
	}

	ctx.JSON(http.StatusOK, weight)
}

func (c *weightController) Update(ctx *gin.Context) {
	var weight models.Weight
	v := weight_validation.NewValidation()

	json := weight_validation.GetUpdateRule{}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": validation.GetError(err.(validator.ValidationErrors), v.GetMessage()),
		})
		return
	}

	if err := models.Db.Model(&models.Weight{}).
		Scopes((&models.Weight{}).GetWeight(ctx.Param("deviceid"), ctx.Param("id"))).
		First(&weight).Error; err != nil {

			ctx.JSON(http.StatusNotFound, gorm.ErrRecordNotFound)
			return
	}

	if err := models.Db.Model(&weight).Update(json).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, weight)
}

func (c *weightController) Store(ctx *gin.Context) {

	v := weight_validation.NewValidation()

	json := weight_validation.GetStoreRule{}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": validation.GetError(err.(validator.ValidationErrors), v.GetMessage()),
		})
		return
	}

	weight := models.Weight{
		Number: json.Number,
	}

	if err := models.Db.Create(&weight).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, weight)
}
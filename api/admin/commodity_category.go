package admin

import (
	"singo/api"
	"singo/service/admin/commodity_category"

	"github.com/gin-gonic/gin"
)

// CreateCommodityCategory 创建
func CreateCommodityCategory(c *gin.Context) {
	var service commodity_category.CreateCommodityCategoryService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}
}

func UpdateCommodityCategory(c *gin.Context) {
	var service commodity_category.CreateCommodityCategoryService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}
}

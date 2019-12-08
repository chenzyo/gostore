package commodity

import (
	"singo/serializer"

	"github.com/gin-gonic/gin"
)

type CommodityCreateService struct {
	CategoryId   string     `form:"category_id" json:"category_id" binding:"required"`
	Name         string     `form:"name" json:"name" binding:"required"`
	SellingPoint string     `form:"selling_point" json:"selling_point" binding:"required"`
	Products     []products `form:"products" json:"products" binding:"required"`
}

type products struct {
	Name  string `form:"name" json:"name" binding:"required"`
	Stock string `form:"name" json:"name" binding:"required"`
	Price string `form:"name" json:"name" binding:"required"`
}

// Login 用户登录函数
func (service *CommodityCreateService) Login(c *gin.Context) serializer.Response {

	data := map[string]string{
		"access_token": "213",
	}
	return serializer.Response{
		Data: data,
	}
}

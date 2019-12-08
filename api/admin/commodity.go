package admin

import (
	"singo/api"
	"singo/service/admin/commodity"

	"github.com/gin-gonic/gin"
)

func CommdityCreate(c *gin.Context) {
	var service commodity.CommodityCreateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}
}

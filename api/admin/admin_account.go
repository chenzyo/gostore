package admin

import (
	"singo/api"
	"singo/service/site/user"

	"github.com/gin-gonic/gin"
)

//AdminLogin 管理员账户登录
func AccountLogin(c *gin.Context) {
	var service user.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, api.ErrorResponse(err))
	}
}

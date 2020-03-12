package account

import (
	"singo/middleware/jwt"
	"singo/model"
	"singo/serializer"

	"github.com/gin-gonic/gin"
)

type AccountLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Login 用户登录函数
func (service *AccountLoginService) Login(c *gin.Context) serializer.Response {
	var user model.AdminAccount

	if err := model.DB.Select("id,user_name,password_digest").Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	token, err := service.generateToken(user)
	if err != nil {
		return serializer.DBErr("token加密失败", nil)
	}

	data := map[string]string{
		"access_token": token,
	}
	return serializer.Response{
		Data: data,
	}
}

// 生成令牌
func (service *AccountLoginService) generateToken(user model.AdminAccount) (string, error) {

	msg := jwt.Msg{
		UserID:   user.ID,
		UserName: user.UserName,
	}

	return jwt.Encode(msg)
}

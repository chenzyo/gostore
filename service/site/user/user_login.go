package user

import (
	"singo/middleware"
	"singo/model"
	"singo/serializer"
	"time"

	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
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
func (service *UserLoginService) generateToken(user model.User) (string, error) {
	j := middleware.Jwt{}

	claims := middleware.Msg{
		ID: user.ID,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "brand",                         //签名的发行者
		},
	}

	return j.CreateToken(claims)
}

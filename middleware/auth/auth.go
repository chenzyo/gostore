package auth

import (
	"singo/middleware/jwt"
	"singo/model"
	"singo/serializer"
	"strings"

	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser(platform string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		kv := strings.Split(token, " ")
		if len(kv) != 2 || kv[0] != "Bearer" {
			c.Next()
		}
		claims := jwt.Msg{}
		// parseToken 解析token包含的信息
		claims, err := jwt.Decode(kv[1])
		if err != nil {
			c.Next()
		}
		//检查是否存在用户，如果存在，保存用户信息至context中
		c.Set("user", nil)
		if claims.UserID != 0 {
			if platform == "site" {
				//用户端
				user, err := model.GetUser(claims.UserID)
				if err == nil {
					c.Set("user", &user)
				}
			} else {
				//后台
				user, err := model.GetAdminUser(claims.UserID)
				if err == nil {
					c.Set("user", &user)
				}
			}
		}
		c.Next()
	}
}

//AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			//if _, ok := user.(*model.User); ok {
			//}
			c.Next()
			return
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}

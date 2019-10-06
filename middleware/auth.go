package middleware

import (
	"fmt"
	"singo/model"
	"singo/serializer"
	"strings"

	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		kv := strings.Split(token, " ")
		if len(kv) != 2 || kv[0] != "Bearer" {
			c.Next()
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(kv[1])
		fmt.Println(claims)
		fmt.Println(token)
		fmt.Println(err)
		if err != nil {
			if err == TokenExpired {
				c.Next()
			}
			c.Next()
		}

		if claims != nil {
			user, err := model.GetUser(claims.ID)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

//AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}

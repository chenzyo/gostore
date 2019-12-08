package server

import (
	"os"
	"singo/api"
	"singo/api/admin"
	"singo/api/site"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", site.UserRegister)

		// 用户登录
		v1.POST("user/login", site.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", site.UserMe)
			auth.DELETE("user/logout", site.UserLogout)
		}
	}

	adm := r.Group("/api/v1/admin")
	{
		adm.POST("commodity/create", admin.CommdityCreate)
	}

	return r
}
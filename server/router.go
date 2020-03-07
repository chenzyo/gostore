package server

import (
	"os"
	"singo/api"
	"singo/api/admin"
	"singo/api/site"
	"singo/middleware/auth"
	"singo/middleware/cors"
	"singo/middleware/session"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	//启用session
	r.Use(session.Session(os.Getenv("SESSION_SECRET")))
	//处理跨域
	r.Use(cors.Cors())

	// 商城前台路由
	v1 := r.Group("/api/v1")
	{
		//获取鉴权会员信息
		v1.Use(auth.CurrentUser("site"))

		//ping
		v1.POST("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", site.UserRegister)

		// 用户登录
		v1.POST("user/login", site.UserLogin)

		// 需要登录保护的
		authRequire := v1.Group("").Use(auth.AuthRequired())
		{
			// User Routing
			authRequire.GET("user/me", site.UserMe)
			authRequire.DELETE("user/logout", site.UserLogout)
		}
	}

	adm := r.Group("/api/v1/admin").Use(auth.AuthRequired())
	{
		adm.Use(auth.CurrentUser("admin"))
		adm.POST("commodity/create", admin.CommodityCreate)
	}

	return r
}

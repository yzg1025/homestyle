package initialization

import (
	"fmt"
	"gin/middleware"
	"gin/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Cors())

	BaseGroup := Router.Group("api")
	{
       router.BaseRouter(BaseGroup)
	}

	PHSGroup := Router.Group("api")
	PHSGroup.Use(middleware.JWTAuth())
	{
		router.PrivacyRouter(PHSGroup)
	}
	fmt.Println("路由注册成功")
	return Router
}
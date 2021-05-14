package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "golang-admin/docs"
	"golang-admin/internal/middleware"
	v1 "golang-admin/internal/routers/api/v1"
	"golang-admin/internal/routers/common"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.Tracing())
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	r.Use(middleware.Cors())
	r.Use(middleware.Translations())

	group := r.Group("/common")
	{
		_common := common.NewCommon()
		group.POST("/wechatlogin", _common.WechatLogin)
		group.GET("/getossaccesstoken", _common.GetOssAccessToken)
	}

	apiV1 := r.Group("/api/v1")
	{
		user := v1.NewUser()
		apiV1.POST("/register", user.Register)
		apiV1.POST("/login", user.Login)
		apiV1.PUT("/users/:id", user.Update)
		apiV1.GET("/users/:id", user.Get)
		apiV1.DELETE("/users/:id", user.Delete)

		news := v1.NewNews()
		apiV1.GET("/news", news.List)
		apiV1.POST("/news", news.Create)

		tag := v1.NewTag()
		apiV1.POST("/tags", tag.Create)
		apiV1.GET("/tags", tag.List)
	}

	return r
}

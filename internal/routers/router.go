package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "golang-admin/docs"
	"golang-admin/internal/middleware"
	v1 "golang-admin/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.Translations())

	apiV1 := r.Group("/api/v1")
	{
		user := v1.NewUser()
		apiV1.POST("/register", user.Register)
		apiV1.POST("/login", user.Login)
		apiV1.GET("/users/:id", user.Get)
		apiV1.DELETE("/users/:id", user.Delete)

		news := v1.NewNews()
		apiV1.GET("/news", middleware.JWT(), news.List)
	}

	return r
}

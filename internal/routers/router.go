package routers

import (
	"github.com/gin-gonic/gin"
	"golang-admin/internal/middleware"
	v1 "golang-admin/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Translations())

	apiV1 := router.Group("/api/v1")
	{
		user := v1.NewUser()
		apiV1.POST("/register", user.Register)
		apiV1.POST("/login", user.Login)
		apiV1.GET("/users/:id", user.Get)
		apiV1.DELETE("/users/:id", user.Delete)
	}

	return router
}

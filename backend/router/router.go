package router

import (
	"backend/controller"
	"backend/router/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.GET("/ping", controller.Ping)

  {
    r.POST("/login", controller.Login)
    r.POST("/register", controller.Register)
  }

	admin := r.Group("/admin")
  admin.Use(middleware.Jwt())
  {
    admin.GET("/user", func(ctx *gin.Context) {
      ctx.JSON(200, gin.H{"status": "ok"})
    })
  }

  user := r.Group("/user")
  user.Use(middleware.Jwt())

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"status": "error", "message": "Not Found"})
	})

	return r
}

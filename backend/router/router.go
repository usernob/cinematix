package router

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
  r := gin.Default()
  r.Use()

  r.GET("/ping", controller.Ping)

  // admin := r.Group("/admin")

  r.NoRoute(func(ctx *gin.Context) {
    ctx.JSON(404, gin.H{"status": "404", "message": "Not Found"})
  })

  return r
}

package router

import (
	"backend/controller"
	filmcontroller "backend/controller/filmController"
	penasanancontroller "backend/controller/penasananController"
	usercontroller "backend/controller/userController"
	"backend/router/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

  r.Static("/storage", "./storage")

	r.GET("/ping", controller.Ping)

  {
    r.POST("/login/user", controller.UserLogin)
    r.POST("/register/user", controller.UserRegister)
    r.GET("/films/tayang", filmcontroller.FilmTayang)
    r.GET("/films/akan-tayang", filmcontroller.FilmAkanTayang)
    r.GET("/films", filmcontroller.FilmList)
    r.GET("/films/:id", filmcontroller.FilmDetailPenyangan)
    r.GET("/pesan/:penayangan_id", penasanancontroller.ShowKursi)
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
  {
    user.GET("/info", usercontroller.GetUserInformation)
  }

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"status": "error", "message": "Not Found"})
	})

	return r
}

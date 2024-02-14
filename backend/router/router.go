package router

import (
	"backend/controller"
	admincontroller "backend/controller/adminController"
	filmcontroller "backend/controller/filmController"
	kursicontroller "backend/controller/kursiController"
	usercontroller "backend/controller/userController"
	"backend/router/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(middleware.DeleteExpPenayangan())

	r.Static("/storage", "./storage")

	r.MaxMultipartMemory = 8 << 20
	r.GET("/ping", controller.Ping)

	{
		r.POST("/login/user", controller.UserLogin)
		r.POST("/register/user", controller.UserRegister)
		r.POST("/login/admin", controller.AdminLogin)
		r.POST("/register/admin", controller.AdminRegister)
		r.GET("/films/tayang", filmcontroller.FilmTayang)
		r.GET("/films/akan-tayang", filmcontroller.FilmAkanTayang)
		r.GET("/films", filmcontroller.FilmList)
		r.GET("/films/:id", filmcontroller.FilmDetail)
		r.GET("/films/:id/:penayangan_id", filmcontroller.FilmDetailPenayangan)
		r.GET("/kursi/:penayangan_id", kursicontroller.ShowKursi)
		r.GET("/kursi/:penayangan_id/:kursi_id", kursicontroller.CheckStatusKursi)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.Jwt())
	{
		admin.GET("/info", admincontroller.GetAdminInformation)
	}

	user := r.Group("/user")
	user.Use(middleware.Jwt())
	user.Use(middleware.DeleteExpTiket())
	{
		user.GET("/info", usercontroller.GetUserInformation)
		user.PUT("/profile/update", usercontroller.UpdateProfile)
		user.POST("/pesanan/add", usercontroller.AddPesanan)
		user.POST("/pesanan/update-pembayaran", usercontroller.UpdatePembayaran)
		user.GET("/pesanan", usercontroller.GetAllPesanan)
		user.GET("/pesanan/:tiket_id", usercontroller.GetTiketDone)
		user.GET("/pesanan/:tiket_id/waiting", usercontroller.GetTiketWaiting)
	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"status": "error", "message": "Not Found"})
	})

	return r
}

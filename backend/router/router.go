package router

import (
	"backend/controller"
	admincontroller "backend/controller/adminController"
	filmcontroller "backend/controller/filmController"
	kursicontroller "backend/controller/kursiController"
	tiketcontroller "backend/controller/tiketController"
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

		r.GET("/films/tayang", filmcontroller.FilmTayang)
		r.GET("/films/akan-tayang", filmcontroller.FilmAkanTayang)
		r.GET("/films/populer", filmcontroller.FilmPopuler)
		r.GET("/films", filmcontroller.FilmList)
		r.GET("/films/search", filmcontroller.FilmSearch)
		r.GET("/films/:id", filmcontroller.FilmDetail)
		r.GET("/films/:id/:penayangan_id", filmcontroller.FilmDetailPenayangan)

		r.GET("penayangan", filmcontroller.GetListPenayangan)
    r.GET("penayangan/:id", filmcontroller.GetPenayanganDetail)

		r.GET("/genre/search", filmcontroller.GenreSearch)
		r.GET("/genre", filmcontroller.GenreList)
		r.GET("/genre/:id", filmcontroller.GenreDetail)

		r.GET("/kursi/:penayangan_id", kursicontroller.ShowKursi)
		r.GET("/kursi/:penayangan_id/:kursi_id", kursicontroller.CheckStatusKursi)
	}

	admin := r.Group("/admin")
	admin.Use(middleware.Jwt())
	{
		admin.GET("/info", admincontroller.GetAdminInformation)
		admin.GET("/all-films", filmcontroller.GetAllFilm)
		admin.GET("/report/thisweek", admincontroller.GetThisWeekReport)

    admin.GET("/tiket/search", tiketcontroller.SearchTiket)
    admin.POST("/tiket/checkin/:tiket_id", tiketcontroller.CheckInTiket)
	}

	superAdmin := admin.Use(middleware.SuperAdminRole())
	{
		superAdmin.POST("/register", controller.AdminRegister)
		superAdmin.GET("/list-admin", admincontroller.AdminList)
		superAdmin.DELETE("/:id", admincontroller.DeleteAdmin)

		superAdmin.PUT("/films/:id", filmcontroller.EditFilm)
		superAdmin.POST("/films", filmcontroller.AddFilm)
		superAdmin.DELETE("/films/:id", filmcontroller.DeleteFilm)

		superAdmin.POST("/penayangan", filmcontroller.AddPenayangan)
		superAdmin.DELETE("/penayangan/:id", filmcontroller.DeletePenayangan)
		superAdmin.PUT("/penayangan/:id", filmcontroller.EditPenayangan)

		superAdmin.POST("/genre", filmcontroller.AddGenre)
		superAdmin.PUT("/genre/:id", filmcontroller.EditGenre)
		superAdmin.DELETE("/genre/:id", filmcontroller.DeleteGenre)
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

package admincontroller

import (
	"backend/controller"
	"backend/model"
	"fmt"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetAdminInformation(c *gin.Context) {
	user := c.MustGet("user").(*model.Admin)
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", user))
}

func GetThisWeekReport(c *gin.Context) {
	res, err := model.GetReport(true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", res))
}

func AdminList(c *gin.Context) {
	res, err := model.GetAdminList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", res))
}

func DeleteAdmin(c *gin.Context) {
	strid := c.Param("id")
	id, convErr := strconv.Atoi(strid)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, convErr.Error(), nil))
		return
	}
	err := model.DeleteAdmin(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

type FilmRequest struct {
	Title        string                `json:"title" form:"title"`
	Rating       float32               `json:"rating" form:"rating"`
	TanggalRilis time.Time             `json:"tanggal_rilis" form:"tanggal_rilis"`
	Sinopsis     string                `json:"sinopsis" form:"sinopsis"`
	Poster       *multipart.FileHeader `json:"poster" form:"poster"`
}

func AddFilm(c *gin.Context) {
	var req FilmRequest
	err := c.ShouldBindWith(&req, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, err.Error(), nil))
		return
	}

	fileSavedPath := path.Join("storage", "image", "poster", fmt.Sprintf("%d-%d.png", rand.Intn(80), time.Now().Unix()))
	if req.Poster != nil {
		if err := c.SaveUploadedFile(req.Poster, fileSavedPath); err != nil {
			c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
			return
		}
	}

	addErr := model.AddFilm(model.Film{
		Title:        req.Title,
		Rating:       req.Rating,
		TanggalRilis: req.TanggalRilis,
		Sinopsis:     req.Sinopsis,
		PosterPath:   fileSavedPath,
	})

	if addErr != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, addErr.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

func EditFilm(c *gin.Context) {
	strid := c.Param("id")
	id, convErr := strconv.Atoi(strid)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, convErr.Error(), nil))
		return
	}

	film, err := model.GetFilmDetail(strid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}

	var req FilmRequest

	bindErr := c.ShouldBindWith(&req, binding.FormMultipart)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, bindErr.Error(), nil))
		return
	}

	fileSavedPath := path.Join("storage", "image", "poster", fmt.Sprintf("%d-%d.png", id, time.Now().Unix()))
	if req.Poster != nil {
		removeErr := os.Remove(film.PosterPath)
		if removeErr != nil {
			c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, removeErr.Error(), nil))
			return
		}
		if err := c.SaveUploadedFile(req.Poster, fileSavedPath); err != nil {
			c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
			return
		}
	}

	editErr := model.EditFilm(film)
	if editErr != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, editErr.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

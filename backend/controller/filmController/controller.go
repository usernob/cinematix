package filmcontroller

import (
	"backend/controller"
	"backend/model"
	"fmt"
	"math/rand"
	"mime/multipart"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func FilmTayang(c *gin.Context) {
	data, err := model.GetFilmTayang(false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", data))
}
func FilmAkanTayang(c *gin.Context) {
	data, err := model.GetFilmTayang(true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", data))
}

func FilmPopuler(c *gin.Context) {
	data, err := model.GetFilmPopuler()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", data))
}

func GetAllFilm(c *gin.Context) {
	res, err := model.GetAllFilm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", res))
}

func FilmList(c *gin.Context) {
	data, err := model.GetFilmHaveTayang()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", data))
}

func FilmDetail(c *gin.Context) {
	strid := c.Param("id")
	id, convErr := strconv.Atoi(strid)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, convErr.Error(), nil))
		return
	}

	data, err := model.GetFilmDetail(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", data))
}

func FilmDetailPenayangan(c *gin.Context) {
	id := c.Param("id")
	penayangan_id := c.Param("penayangan_id")
	data, err := model.GetPenayanganSingle(id, penayangan_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", data))
}

type FilmRequest struct {
	Title        string                `form:"title" binding:"required"`
	Rating       float32               `form:"rating" binding:"required"`
	TanggalRilis time.Time             `form:"tanggal_rilis" time_format:"2006-01-02" binding:"required"`
	Sinopsis     string                `form:"sinopsis" binding:"required"`
	Genres       []string              `form:"genres[]"`
	Poster       *multipart.FileHeader `form:"poster"`
}

func AddFilm(c *gin.Context) {
	var req FilmRequest
	err := c.ShouldBindWith(&req, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, err.Error(), nil))
		return
	}

	if req.Poster == nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, "Poster is required", nil))
		return
	}
	fileSavedPath := path.Join("storage", "image", "poster", fmt.Sprintf("%d-%d.png", rand.Intn(80), time.Now().Unix()))

	if err := c.SaveUploadedFile(req.Poster, fileSavedPath); err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}

	var genres []int
	if req.Genres != nil {
		for _, genre := range req.Genres {
			genreId, err := strconv.Atoi(genre)
			if err != nil {
				c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
				return
			}
			genres = append(genres, genreId)
		}
	}

	addErr := model.AddFilm(req.Title, req.Rating, req.TanggalRilis, req.Sinopsis, fileSavedPath, genres)

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

	var req FilmRequest

	bindErr := c.ShouldBindWith(&req, binding.FormMultipart)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, bindErr.Error(), nil))
		return
	}

	var fileSavedPath string
	if req.Poster != nil {
		fileSavedPath = path.Join("storage", "image", "poster", fmt.Sprintf("%d-%d.png", id, time.Now().Unix()))
		err := c.SaveUploadedFile(req.Poster, fileSavedPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
			return
		}
	}

	var genres []int
	if req.Genres != nil {
		for _, genre := range req.Genres {
			genreId, err := strconv.Atoi(genre)
			if err != nil {
				c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
				return
			}
			genres = append(genres, genreId)
		}
	}

	editErr := model.EditFilm(id, req.Title, req.Rating, req.TanggalRilis, req.Sinopsis, fileSavedPath, genres)
	if editErr != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, editErr.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

func DeleteFilm(c *gin.Context) {
	strid := c.Param("id")
	id, convErr := strconv.Atoi(strid)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, convErr.Error(), nil))
		return
	}
	err := model.DeleteFilm(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

func GenreSearch(c *gin.Context) {
	genre := c.Query("q")
	data, err := model.FindGenres(genre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", data))
}

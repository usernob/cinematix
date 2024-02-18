package filmcontroller

import (
	"backend/controller"
	"backend/model"
	"backend/pkg/logjson"
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
	id := c.Param("id")
	data, err := model.GetFilmDetail(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	logjson.ToJSON(data)
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
	Title        string                `form:"title"`
	Rating       float32               `form:"rating"`
	TanggalRilis time.Time             `form:"tanggal_rilis" time_format:"2006-01-02"`
	Sinopsis     string                `form:"sinopsis"`
	Poster       *multipart.FileHeader `form:"poster"`
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
    film.PosterPath = fileSavedPath
	}
  film.Title = req.Title
  film.Rating = req.Rating
  film.TanggalRilis = req.TanggalRilis
  film.Sinopsis = req.Sinopsis

	editErr := model.EditFilm(uint(id), film)
	if editErr != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, editErr.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", film))
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

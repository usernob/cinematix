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

func FilmSearch(c *gin.Context) {
  q := c.Query("q")
  data, err := model.FindFilm(q)
  if err != nil {
    c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
    return
  }
  c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", data))
}

func GetListPenayangan(c *gin.Context) {
	data, err := model.GetListPenayangan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", data))
}

func GetPenayanganDetail(c *gin.Context) {
  strid := c.Param("id")
  id, convErr := strconv.Atoi(strid)
  if convErr != nil {
    c.JSON(http.StatusBadRequest, controller.Response(controller.Error, convErr.Error(), nil))
    return
  }
  data, err := model.GetPenayanganDetail(uint(id))
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

func GenreList(c *gin.Context) {
	data, err := model.GenreList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", data))
}

func GenreDetail(c *gin.Context) {
	strid := c.Param("id")
	id, convErr := strconv.Atoi(strid)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, convErr.Error(), nil))
		return
	}
	data, err := model.GenreDetail(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", data))
}

type GenreRequest struct {
	Nama string `json:"nama" form:"nama" binding:"required"`
}

func AddGenre(c *gin.Context) {
	var req GenreRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	addErr := model.AddGenre(req.Nama)
	if addErr != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, addErr.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

func EditGenre(c *gin.Context) {
	var req GenreRequest
	strid := c.Param("id")
	id, convErr := strconv.Atoi(strid)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, convErr.Error(), nil))
		return
	}

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, err.Error(), nil))
		return
	}

	editErr := model.EditGenre(uint(id), req.Nama)
	if editErr != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, editErr.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

func DeleteGenre(c *gin.Context) {
	strid := c.Param("id")
	id, convErr := strconv.Atoi(strid)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, convErr.Error(), nil))
		return
	}

	err := model.DeleteGenre(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

type PenayanganRequest struct {
	FilmID        uint      `json:"film_id" form:"film_id" binding:"required"`
	AudiotoriumID uint      `json:"audiotorium_id" form:"audiotorium_id" binding:"required"`
	Tanggal       time.Time `json:"tanggal" form:"tanggal" binding:"required" time_format:"2006-01-02"`
	Mulai         time.Time `json:"mulai" form:"mulai" binding:"required" time_format:"15:04"`
	Selesai       time.Time `json:"selesai" form:"selesai" binding:"required" time_format:"15:04"`
	Harga         uint      `json:"harga" form:"harga" binding:"required"`
}

func fixDateTime(date time.Time, datetime time.Time) time.Time {
  return time.Date(date.Year(), date.Month(), date.Day(), datetime.Hour(), datetime.Minute(), 0, 0, date.Location())
}

func AddPenayangan(c *gin.Context) {
  var req PenayanganRequest
  bindErr := c.ShouldBind(&req)
  if bindErr != nil {
    c.JSON(http.StatusBadRequest, controller.Response(controller.Error, bindErr.Error(), nil))
    return
  }
  err := model.AddPenayangan(req.FilmID, req.Tanggal, fixDateTime(req.Tanggal, req.Mulai), fixDateTime(req.Tanggal, req.Selesai), req.AudiotoriumID, req.Harga)
  if err != nil {
    c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
    return
  }
  c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

func EditPenayangan(c *gin.Context) {
  var req PenayanganRequest
  strid := c.Param("id")
  id, convErr := strconv.Atoi(strid)
  if convErr != nil {
    c.JSON(http.StatusBadRequest, controller.Response(controller.Error, convErr.Error(), nil))
    return
  }
  bindErr := c.ShouldBind(&req)
  if bindErr != nil {
    c.JSON(http.StatusBadRequest, controller.Response(controller.Error, bindErr.Error(), nil))
    return
  }
  err := model.UpdatePenayangan(uint(id), req.FilmID, req.Tanggal, fixDateTime(req.Tanggal, req.Mulai), fixDateTime(req.Tanggal, req.Selesai), req.AudiotoriumID, req.Harga)
  if err != nil {
    c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
    return
  }
  c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

func DeletePenayangan(c *gin.Context) {
  strid := c.Param("id")
  id, convErr := strconv.Atoi(strid)
  if convErr != nil {
    c.JSON(http.StatusBadRequest, controller.Response(controller.Error, convErr.Error(), nil))
    return
  }
  err := model.DeletePenayangan(uint(id))
  if err != nil {
    c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
    return
  }

  c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

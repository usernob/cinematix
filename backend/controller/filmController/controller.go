package filmcontroller

import (
	"backend/controller"
	"backend/model"
	"backend/pkg/logjson"
	"net/http"

	"github.com/gin-gonic/gin"
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

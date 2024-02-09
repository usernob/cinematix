package filmcontroller

import (
	"backend/controller"
	"backend/model"
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

package tiketcontroller

import (
	"backend/controller"
	"backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SearchTiket(c *gin.Context) {
	str_user_id := c.Query("user_id")
	str_tiket_id := c.Query("tiket_id")
	user_id, err := strconv.ParseUint(str_user_id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, "user_id is not number", nil))
		return
	}

	tiket_id, err := strconv.ParseUint(str_tiket_id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, "tiket_id is not number", nil))
		return
	}

	data, err := model.GetPesanan(uint(tiket_id), uint(user_id), model.Done)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", data))
}

func CheckInTiket(c *gin.Context) {
	admin := c.MustGet("user").(*model.Admin)
  strtiket_id := c.Param("tiket_id")
  tiket_id, err := strconv.ParseUint(strtiket_id, 10, 32)
  if err != nil {
    c.JSON(http.StatusBadRequest, controller.Response(controller.Error, "tiket_id is not number", nil))
    return
  }
	cErr := model.CheckInTiket(uint(tiket_id), admin.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, cErr.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

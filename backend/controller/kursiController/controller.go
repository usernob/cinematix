package kursicontroller

import (
	"backend/controller"
	"backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func ShowKursi(c *gin.Context) {
  intPenyanganID, err := strconv.ParseUint(c.Param("penayangan_id"), 10, 32)
  if err != nil {
    c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, "penayangan_id is not a number", nil))
    return
  }
  data, err := model.GetKursi(uint(intPenyanganID))
  if err != nil {
    c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
    return
  }
  c.JSON(http.StatusOK, controller.Response(controller.Ok, "success", data))
}

func CheckStatusKursi(c *gin.Context) {
  
  intPenyanganID, err := strconv.ParseUint(c.Param("penayangan_id"), 10, 32)
  if err != nil {
    c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, "penayangan_id is not a number", nil))
    return
  }
  intKursiID, err := strconv.ParseUint(c.Param("kursi_id"), 10, 32)
  if err != nil {
    c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, "kursi_id is not a number", nil))
    return
  }
  data, err := model.GetKursi(uint(intPenyanganID))
  if err != nil {
    c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
    return
  }

  for _, value := range data {
    if len(value.Seat) <= 0 {
      continue
    }

    if value.ID == uint(intKursiID) {
      c.JSON(http.StatusConflict, controller.Response(controller.Error, "kursi sudah terisi", nil))
      return
    }
  }

  c.JSON(http.StatusOK, controller.Response(controller.Ok, "succes", nil))
}

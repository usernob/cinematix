package usercontroller

import (
	"backend/controller"
	"backend/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserInformation(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", user))
}

type PesananRequest struct {
  PenayanganId uint `json:"penayangan_id"`
  Kursis []*model.Kursi `json:"kursis"`
  FilmId uint `json:"film_id"`
  TotalHarga uint `json:"total_harga"`
}

func AddPesanan(c *gin.Context) {
  user := c.MustGet("user").(*model.User)

  var formData PesananRequest
  err := c.ShouldBindJSON(&formData)
  if err != nil {
    fmt.Println(err)
    c.JSON(http.StatusBadRequest, controller.Response(controller.Error, err.Error(), nil))
    return
  }
  tiket := &model.Tiket{
    UserID: user.ID,
    PenayanganID: formData.PenayanganId,
    TotalHarga: formData.TotalHarga,
  }


  err = model.CreateTiketAndSeat(tiket, formData.Kursis)
  if err != nil {
    c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
    return
  }
  c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", tiket))
}

func UpdatePembayaran(c *gin.Context) {
  
}

func GetPesanan(c *gin.Context) {
  user := c.MustGet("user").(*model.User)
  pesanan_id, err := strconv.ParseUint(c.Param("tiket_id"), 10, 32)
  if err != nil {
    c.JSON(http.StatusBadRequest, controller.Response(controller.Error, "pesanan_id is not number", nil))
    return
  }

  res, err := model.GetPesanan(uint(pesanan_id), user.ID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
    return
  }
  c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", res))
}

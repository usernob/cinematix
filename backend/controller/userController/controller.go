package usercontroller

import (
	"backend/controller"
	"backend/model"
	"backend/pkg/logjson"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserInformation(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", user))
}

type PesananRequest struct {
	PenayanganId uint           `json:"penayangan_id"`
	Kursis       []*model.Kursi `json:"kursis"`
	FilmId       uint           `json:"film_id"`
	TotalHarga   uint           `json:"total_harga"`
}

type UpdateProfileRequest struct {
	Email string `json:"email"`
	Nama  string `json:"nama"`
}

func UpdateProfile(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	fmt.Println(c.GetHeader("Content-Type"))
	var userUpdate UpdateProfileRequest
	err := c.ShouldBindJSON(&userUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, err.Error(), nil))
		return
	}

	user.Email = userUpdate.Email
	user.Nama = userUpdate.Nama
	fileSavedPath := path.Join("storage", "images", "profile", fmt.Sprintf("%d.png", user.ID))

	file, err := c.FormFile("avatar")
  fmt.Println(err)
	if err == nil {
		if err := c.SaveUploadedFile(file, fileSavedPath); err != nil {
			c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
			return
		}
		user.Avatar = &fileSavedPath
	}

	logjson.ToJSON(userUpdate)
	logjson.ToJSON(user)
	Updateerr := model.UpdateUser(user)
	if Updateerr != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, Updateerr.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", user))
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
		UserID:       user.ID,
		PenayanganID: formData.PenayanganId,
		TotalHarga:   formData.TotalHarga,
	}

	err = model.CreateTiketAndSeat(tiket, formData.Kursis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", tiket))
}

func UpdatePembayaran(c *gin.Context) {
	var tiket model.Tiket
	err := c.ShouldBindJSON(&tiket)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, err.Error(), nil))
		return
	}

	Updateerr := model.UpdateTiketPembayaran(&tiket)
	if Updateerr != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, Updateerr.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", tiket))
}



func GetAllPesanan(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	res, err := model.GetAllTiket(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", res))
}

func GetTiketDone(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	pesanan_id, err := strconv.ParseUint(c.Param("tiket_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, "tiket_id is not number", nil))
		return
	}

	res, err := model.GetPesanan(uint(pesanan_id), user.ID, model.Done)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", res))
}


func GetTiketWaiting(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	pesanan_id, err := strconv.ParseUint(c.Param("tiket_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, "tiket_id is not number", nil))
		return
	}

	res, err := model.GetPesanan(uint(pesanan_id), user.ID, model.Waiting)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", res))
}

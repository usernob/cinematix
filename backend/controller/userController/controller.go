package usercontroller

import (
	"backend/controller"
	"backend/model"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	Email  string                `form:"email" binding:"required"`
	Nama   string                `form:"nama" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar"`
}

func UpdateProfile(c *gin.Context) {
	user := c.MustGet("user").(*model.User)

	var userUpdate UpdateProfileRequest
	err := c.ShouldBindWith(&userUpdate, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, err.Error(), nil))
		return
	}

	user.Email = userUpdate.Email
	user.Nama = userUpdate.Nama
	fileSavedPath := path.Join("storage", "image", "profile", fmt.Sprintf("%d%d.png", time.Now().Unix(), user.ID))

	if userUpdate.Avatar != nil {
		if user.Avatar != nil {
			removeErr := os.Remove(*user.Avatar)
			if removeErr != nil {
				c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, removeErr.Error(), nil))
				return
			}
		}

		if err := c.SaveUploadedFile(userUpdate.Avatar, fileSavedPath); err != nil {
			c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
			return
		}
		user.Avatar = &fileSavedPath
	}

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

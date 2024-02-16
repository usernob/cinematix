package admincontroller

import (
	"backend/controller"
	"backend/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAdminInformation(c *gin.Context) {
	user := c.MustGet("user").(*model.Admin)
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", user))
}

func GetThisWeekReport(c *gin.Context) {
	res, err := model.GetReport(true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", res))
}

func AdminList(c *gin.Context) {
	res, err := model.GetAdminList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", res))
}

func DeleteAdmin(c *gin.Context) {
	strid := c.Param("id")
	id, convErr := strconv.Atoi(strid)
	if convErr != nil {
		c.JSON(http.StatusBadRequest, controller.Response(controller.Error, convErr.Error(), nil))
		return
	}
	err := model.DeleteAdmin(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, controller.Response(controller.Error, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", nil))
}

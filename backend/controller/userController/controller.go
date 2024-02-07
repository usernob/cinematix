package usercontroller

import (
	"backend/controller"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInformation(c *gin.Context) {
	user := c.MustGet("user").(*model.User)
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", user))
}

package admincontroller

import (
	"backend/controller"
	"backend/model"
	"backend/pkg/logjson"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetAdminInformation(c *gin.Context) {
	user := c.MustGet("user").(*model.Admin)
  logjson.ToJSON(user)
	c.JSON(http.StatusOK, controller.Response(controller.Ok, "Success", user))
}

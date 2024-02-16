package middleware

import (
	"backend/controller"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuperAdminRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.MustGet("user").(*model.Admin).Role
		if role != model.SuperAdminRole {
			c.JSON(http.StatusUnauthorized, controller.Response(controller.Error, "Unauthorized", nil))
			c.Abort()
			return
		}
		c.Next()
	}
}

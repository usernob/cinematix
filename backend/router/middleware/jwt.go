package middleware

import (
	"backend/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"backend/pkg/jwt"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {

		headerToken := c.GetHeader("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Bearer token not found", "data": nil})
			c.Abort()
			return
		}

		tokenString := strings.Replace(headerToken, "Bearer ", "", -1)

		token, err := jwt.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Token is not valid", "data": nil})
			c.Abort()
			return
		}

		user_id, is_admin, err := jwt.GetUserIdFromToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized", "data": nil})
			c.Abort()
			return
		}

		if is_admin {
			admin, err := model.GetUserBy(model.User{ID: user_id})
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized", "data": nil})
				c.Abort()
				return
			}
			c.Set("user", admin)
		} else {
			user, err := model.GetUserBy(model.User{ID: user_id})
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized", "data": nil})
				c.Abort()
				return
			}
			c.Set("user", user)
		}

		c.Set("is_admin", is_admin)

		c.Next()
	}
}

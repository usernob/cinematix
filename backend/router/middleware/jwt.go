package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"backend/pkg/jwt"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {

		headerToken := c.GetHeader("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Bearer token not found"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(headerToken, "Bearer ", "", -1)

		token, err := jwt.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Token is not valid"})
			c.Abort()
			return
		}

		user_id, is_admin, err := jwt.GetUserIdFromToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user_id", user_id)
		c.Set("is_admin", is_admin)

		c.Next()
	}
}

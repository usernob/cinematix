package middleware

import (
	"backend/model"

	"github.com/gin-gonic/gin"
)

func DeleteExpPenayangan() gin.HandlerFunc {
	return func(c *gin.Context) {
		model.DeleteExpPenayangan()
		c.Next()
	}
}

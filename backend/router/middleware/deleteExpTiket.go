package middleware

import (
	"backend/model"

	"github.com/gin-gonic/gin"
)

func DeleteExpTiket() gin.HandlerFunc {
  return func(c *gin.Context) {
    model.DeleteExpTiket()
    c.Next()
  }
}

package controller

import "github.com/gin-gonic/gin"

const (
	Ok    string = "ok"
	Error string = "error"
)

func Response(status string, message string, data interface{}) gin.H {
	return gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	}
}

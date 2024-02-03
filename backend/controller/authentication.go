package controller

import (
	"backend/model"
	"backend/pkg/jwt"
	"backend/pkg/passwordhash"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}


const (
  Ok string = "ok"
  Error string = "error"
)

func Response(status string, message string, data interface{}) gin.H {
  return gin.H{
    "status":  status,
    "message": message,
    "data":    data,
  }
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, Response(Error, err.Error(), nil))
		return
	}

	user, err := model.GetUserByEmail(json.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, Response(Error, err.Error(), nil))
		return
	}

	if !passwordhash.CheckPasswordHash(json.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, Response(Error, "Wrong email or password", nil))
		return
	}

	token, err := jwt.GenerateToken(user.ID, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response(Error, err.Error(), nil))
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
  c.JSON(http.StatusOK, Response(Ok, "Success login", map[string]string{"token": token}))
}

type RegisterRequest struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, Response(Error, err.Error(), nil))
		return
	}

	passwordHased, err := passwordhash.HashPassword(json.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response(Error, err.Error(), nil))
		return
	}

	json.Password = passwordHased

	user, insErr := model.CreateUser(json.Nama, json.Email, json.Password)
	if insErr != nil {
		c.JSON(http.StatusInternalServerError, Response(Error, insErr.Error(), nil))
		return
	}

	token, err := jwt.GenerateToken(user.ID, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response(Error, err.Error(), nil))
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, Response(Ok, "Success register", map[string]string{"token": token}))
}

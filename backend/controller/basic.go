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

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	user, err := model.GetUserByEmail(json.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Email not found"})
		return
	}

	if !passwordhash.CheckPasswordHash(json.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Wrong password"})
		return
	}

	token, err := jwt.GenerateToken(user.ID, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": token})
}

type RegisterRequest struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	passwordHased, err := passwordhash.HashPassword(json.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	json.Password = passwordHased

	user, insErr := model.CreateUser(json.Nama, json.Email, json.Password)
	if insErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	token, err := jwt.GenerateToken(user.ID, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": token})
}

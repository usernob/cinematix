package controller

import (
	"backend/model"
	"backend/pkg/jwt"
	"backend/pkg/passwordhash"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type ResponseUserAuth struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserLogin(c *gin.Context) {
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
	c.JSON(http.StatusOK, Response(Ok, "Success login", ResponseUserAuth{Token: token, User: *user}))
}

type RegisterRequest struct {
	Nama string `json:"nama" binding:"required"`
	LoginRequest
}

func UserRegister(c *gin.Context) {
	var json RegisterRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, Response(Error, err.Error(), nil))
		return
	}

	_, err := model.GetUserByEmail(json.Email)
  if !(errors.Is(err, gorm.ErrRecordNotFound)) {
    c.JSON(http.StatusBadRequest, Response(Error, "email already exist", nil))
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
	c.JSON(http.StatusOK, Response(Ok, "Success register", ResponseUserAuth{Token: token, User: *user}))
}



type ResponseAdminAuth struct {
	Token string      `json:"token"`
	Admin model.Admin `json:"user"`
}

func AdminLogin(c *gin.Context) {
	var json LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, Response(Error, err.Error(), nil))
		return
	}

	admin, err := model.GetAdminByEmail(json.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, Response(Error, err.Error(), nil))
		return
	}

	if !passwordhash.CheckPasswordHash(json.Password, admin.Password) {
		c.JSON(http.StatusUnauthorized, Response(Error, "Wrong email or password", nil))
		return
	}

	token, err := jwt.GenerateToken(admin.ID, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response(Error, err.Error(), nil))
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, Response(Ok, "Success login", ResponseAdminAuth{Token: token, Admin: *admin}))
}

func AdminRegister(c *gin.Context) {
	var json RegisterRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, Response(Error, err.Error(), nil))
		return
	}

	_, err := model.GetAdminByEmail(json.Email)
  if !(errors.Is(err, gorm.ErrRecordNotFound)) {
    c.JSON(http.StatusBadRequest, Response(Error, "email already exist", nil))
    return
  }

	passwordHased, err := passwordhash.HashPassword(json.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response(Error, err.Error(), nil))
		return
	}

	json.Password = passwordHased

	admin, insErr := model.CreateAdmin(json.Nama, json.Email, json.Password)
	if insErr != nil {
		c.JSON(http.StatusInternalServerError, Response(Error, insErr.Error(), nil))
		return
	}

	token, err := jwt.GenerateToken(admin.ID, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response(Error, err.Error(), nil))
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, Response(Ok, "Success register", ResponseAdminAuth{Token: token, Admin: *admin}))
}

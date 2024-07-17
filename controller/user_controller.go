package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"task/models"
)

type UserController struct {
	DB *gorm.DB
}

func (u *UserController) Login(c *gin.Context) {
	user := models.User{}
	errBindJson := c.ShouldBindJSON(&user)
	if errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errBindJson.Error()})
		return
	}

	password := user.Password

	errDB := u.DB.Where("email = ?", user.Email).Take(&user).Error
	if errDB != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email or Password invalid"})
		return
	}

	errHash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errHash != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or Password invalid"})
		return
	}
	c.JSON(http.StatusOK, user)
}

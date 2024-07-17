package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task/config"
	"task/controller"
	"task/models"
)

func main() {
	//Database
	db := config.DatabaseConnection()
	db.AutoMigrate(&models.User{}, &models.Task{})
	config.CreateOwnerAccount(db)

	//controller
	userController := controller.UserController{DB: db}
	//taskController := controller.TaskController{DB: db}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to the task manager")
	})

	router.POST("/users/login", userController.Login)
	router.Static("/attachment", "./attachment")
	router.Run("192.168.186.205:8080")
}

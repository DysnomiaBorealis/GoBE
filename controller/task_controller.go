package controller

import "gorm.io/gorm"

type TaskController struct {
	DB *gorm.DB
}

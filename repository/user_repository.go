package repository

import (
	"github.com/alejandroas97/online-streaming/db"
	"github.com/alejandroas97/online-streaming/models"
)

func GetAllUsers() *[]models.User {
	var users []models.User

	if err := db.DB.Find(&users).Error; err != nil {
		return nil
	}

	return &users
}

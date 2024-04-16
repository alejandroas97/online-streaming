package repository

import (
	"fmt"

	"github.com/alejandroas97/online-streaming/db"
	"github.com/alejandroas97/online-streaming/models"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers() *[]models.User {
	var users []models.User

	if err := db.DB.Find(&users).Error; err != nil {
		return nil
	}

	for i := range users {
		users[i].Password = ""
	}

	return &users
}

func CreateUser(user *models.User) (*models.User, error) {
	userLogged := GetUserByUsername(user)

	if userLogged != nil {
		return nil, fmt.Errorf("el usuario %s ya está creado", user.Username)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	if err := db.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByUsername(user *models.User) *models.User {
	if err := db.DB.Where("username = ?", user.Username).First(&user).Error; err != nil {
		return nil
	}
	return user
}

func GetUserById(user *models.User) *models.User {
	if err := db.DB.Where("id = ?", user.ID).First(&user).Error; err != nil {
		return nil
	}

	user.Password = ""
	return user
}

package repository

import (
	"fmt"

	"github.com/alejandroas97/online-streaming/db"
	"github.com/alejandroas97/online-streaming/models"
)

func CreateFilm(film *models.Film) (*models.Film, error) {
	foundFilm := FindFilm(film)

	if foundFilm != nil {
		return nil, fmt.Errorf("la película %s ya existe en la base de datos", film.Title)
	}

	if err := db.DB.Create(&film).Error; err != nil {
		return nil, err
	}
	return film, nil
}

func FindFilm(film *models.Film) *models.Film {
	if err := db.DB.Where("title = ?", film.Title).First(&film).Error; err != nil {
		return nil
	}
	return film
}

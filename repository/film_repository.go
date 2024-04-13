package repository

import (
	"gorm.io/gorm"

	"github.com/alejandroas97/online-streaming/models"
)

type FilmRepository struct {
	db *gorm.DB
}

func (r *FilmRepository) GetFilmByID(filmID int) (*models.Film, error) {
	var film models.Film
	film.ID = uint(filmID)

	if err := r.db.First(&film).Error; err != nil {
		return nil, err
	}

	return &film, nil
}

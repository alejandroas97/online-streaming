package repository

import (
	"fmt"

	"github.com/alejandroas97/online-streaming/db"
	"github.com/alejandroas97/online-streaming/models"
)

func GetAllFilms() *[]models.Film {
	var films []models.Film

	if err := db.DB.Find(&films).Error; err != nil {
		return nil
	}

	var user models.User

	for i := range films {
		(films)[i].CreateUser = user
	}

	return &films
}

func CreateFilm(film *models.Film) (*models.Film, error) {
	foundFilm := GetFilmByName(film)

	if foundFilm != nil {
		return nil, fmt.Errorf("la película %s ya existe en la base de datos", film.Title)
	}

	if err := db.DB.Create(&film).Error; err != nil {
		return nil, err
	}
	return film, nil
}

func UpdateFilm(film *models.Film) (*models.Film, error) {

	fmt.Println(film)
	if err := db.DB.Save(film).Error; err != nil {
		return nil, err
	}

	return film, nil
}

func DeleteFilm(film *models.Film) (*models.Film, error) {

	if err := db.DB.Delete(film).Error; err != nil {
		return nil, err
	}

	return film, nil
}

func GetFilmByName(film *models.Film) *models.Film {
	if err := db.DB.Where("title = ?", film.Title).First(&film).Error; err != nil {
		return nil
	}
	return film
}

func GetFilmById(film *models.Film) *models.Film {
	if err := db.DB.Where("id = ?", film.ID).First(&film).Error; err != nil {
		return nil
	}
	return film
}

func GetFilteredFilms(filters map[string]string) (*[]models.Film, error) {
	query := db.DB

	for field, value := range filters {
		switch field {
		case "title":
			query = query.Where("title = ?", value)
		case "releaseDate":
			query = query.Where("release_date = ?", value)
		case "genre":
			query = query.Where("genre = ?", value)
		case "director":
			query = query.Where("director = ?", value)
		default:
			return nil, fmt.Errorf("campo de filtro no válido: %s", field)
		}
	}
	var films []models.Film

	if err := query.Find(&films).Error; err != nil {
		return nil, nil
	}

	return &films, nil
}

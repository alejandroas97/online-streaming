package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/alejandroas97/online-streaming/models"
	"github.com/alejandroas97/online-streaming/repository"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type QueryParams struct {
	Title       string `json:"title,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	Genre       string `json:"genre,omitempty"`
}

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	var film models.Film
	json.NewDecoder(r.Body).Decode(&film)

	tokenString := r.Header.Get("Authorization")

	username, err := getUsernameFromToken(tokenString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	user := &models.User{
		Username: username,
	}

	user = repository.GetUserByUsername(user)

	film.CreateUser = *user

	createdFilm, err := repository.CreateFilm(&film)

	if createdFilm == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userFilm := &models.User{
		Username: username,
	}

	createdFilm.CreateUser = *userFilm

	json.NewEncoder(w).Encode(&film)
}

func GetAllFilms(w http.ResponseWriter, r *http.Request) {
	filters, err := getFilters(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var films *[]models.Film
	if len(filters) > 0 {
		films, err = repository.GetFilteredFilms(filters)
		if err != nil {
			http.Error(w, "Error al obtener las películas filtradas", http.StatusInternalServerError)
			return
		}
	} else {
		films = repository.GetAllFilms()
		if films == nil {
			http.Error(w, "Error al obtener todas las películas", http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(films)
}

func getUsernameFromToken(tokenString string) (string, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	username, ok := token.Claims.(jwt.MapClaims)["username"].(string)
	if !ok {
		return "", fmt.Errorf("nombre de usuario no encontrado en el token")
	}

	return username, nil
}

func getFilters(r *http.Request) (map[string]string, error) {
	query := r.URL.Query()

	filters := make(map[string]string)

	for key, values := range query {
		filters[key] = values[0]
	}

	return filters, nil
}

func GetFilmByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title, ok := vars["title"]
	if !ok {
		http.Error(w, "título de película no especificado", http.StatusBadRequest)
		return
	}

	film := &models.Film{
		Title: title,
	}

	film = repository.GetFilmByName(film)
	if film == nil {
		http.Error(w, "Película no encontrada", http.StatusNotFound)
		return
	}

	user := &models.User{
		ID: film.CreateUserID,
	}

	user = repository.GetUserById(user)

	film.CreateUser = *user

	json.NewEncoder(w).Encode(film)
}

func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	var newFilm *models.Film
	json.NewDecoder(r.Body).Decode(&newFilm)

	vars := mux.Vars(r)
	title, ok := vars["title"]
	if !ok {
		http.Error(w, "título de película no especificado", http.StatusBadRequest)
		return
	}

	film := &models.Film{
		Title: title,
	}

	film = repository.GetFilmByName(film)
	if film == nil {
		http.Error(w, "Película no encontrada", http.StatusNotFound)
		return
	}

	tokenString := r.Header.Get("Authorization")

	username, err := getUsernameFromToken(tokenString)
	if err != nil {
		http.Error(w, "No ha sido posible obtener el usuario", http.StatusNotFound)
		return
	}

	user := &models.User{
		Username: username,
	}

	user = repository.GetUserByUsername(user)

	if film.CreateUserID != user.ID {
		http.Error(w, "Este usuario no puede modificar esta película", http.StatusUnauthorized)
		return
	}

	newFilm.CreateUserID = user.ID
	newFilm.ID = film.ID

	createdFilm, err := repository.UpdateFilm(newFilm)

	if createdFilm == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&createdFilm)
}

func DeleteFilm(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "id de película no especificado", http.StatusBadRequest)
		return
	}

	numId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return
	}

	film := &models.Film{
		ID: uint(numId),
	}

	film = repository.GetFilmById(film)
	if film == nil {
		http.Error(w, "Película no encontrada", http.StatusNotFound)
		return
	}

	tokenString := r.Header.Get("Authorization")

	username, err := getUsernameFromToken(tokenString)
	if err != nil {
		http.Error(w, "No ha sido posible obtener el usuario", http.StatusNotFound)
		return
	}

	user := &models.User{
		Username: username,
	}

	user = repository.GetUserByUsername(user)

	if film.CreateUserID != user.ID {
		http.Error(w, "Este usuario no puede eliminar esta película", http.StatusUnauthorized)
		return
	}

	deletetedFilm, err := repository.DeleteFilm(film)

	if deletetedFilm == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&deletetedFilm)
}

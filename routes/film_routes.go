package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/alejandroas97/online-streaming/models"
	"github.com/alejandroas97/online-streaming/repository"
	"github.com/golang-jwt/jwt"
)

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

	user = repository.FindUser(user)

	film.CreateUser = *user

	createdFilm, err := repository.CreateFilm(&film)

	if createdFilm == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&film)
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

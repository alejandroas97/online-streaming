package routes

import (
	"encoding/json"
	"net/http"

	"github.com/alejandroas97/online-streaming/models"
	"github.com/alejandroas97/online-streaming/repository"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users *[]models.User = repository.GetAllUsers()
	json.NewEncoder(w).Encode(&users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// var
}

package main

import (
	"log"
	"net/http"

	"github.com/alejandroas97/online-streaming/db"
	"github.com/alejandroas97/online-streaming/models"
	"github.com/alejandroas97/online-streaming/routes"
	"github.com/gorilla/mux"
)

const BaseURI = "/online-streaming/v1"

func main() {
	r := mux.NewRouter()

	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Film{})

	r.HandleFunc(BaseURI+"/users", routes.GetAllUsers).Methods("GET")
	r.HandleFunc(BaseURI+"/users", routes.CreateUser).Methods("POST")
	r.HandleFunc(BaseURI+"/users/login", routes.Login).Methods("POST")

	r.HandleFunc(BaseURI+"/films", routes.CheckAuth(routes.CreateFilm)).Methods("POST")

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

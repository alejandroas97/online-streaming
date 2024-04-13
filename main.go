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

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

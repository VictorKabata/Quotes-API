package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/VictorKabata/quotes-api/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) InitializeServer() {
	var err error

	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	DBurl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=diasble", host, port, user, password, dbname)
	server.DB, err = gorm.Open(dbdriver, DBurl)
	if err != nil {
		fmt.Printf("Error connecting to database: %s\n", err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	defer server.DB.Close()

	server.DB.Debug().AutoMigrate(&models.User{})

	server.Router = mux.NewRouter()

	server.InitializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Printf("Listening and serving on port: %s", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	repository "../repositories"
)

type Server struct {
	DB *gorm.DB
	Router   *mux.Router
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	server.DB, err = gorm.Open(DbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DbDriver)
		log.Fatal("Error:", err)
	} else {
		fmt.Printf("Successfully connected to %s database", DbDriver)
	}
	server.DB.Debug().AutoMigrate(&repository.Brew{}, &repository.Respiration{}, &repository.Temperature{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run() {
	log.Fatal(http.ListenAndServe(":8080", server.Router))
}
package api

import (
	"./controllers"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error when loading .env file %v", err)
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST") , os.Getenv("DB_NAME"))

	server.Run()
}
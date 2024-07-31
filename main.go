package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
}

func main() {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error Loading .env File")
	}

	dbUserName := os.Getenv("DB_USERNAME")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASSWORD")

	//connect to a PostgreSQL database
	// Replace the connection details (user, dbname, password, host) with your own
	db, err := sqlx.Connect("postgres", "port="+ dbPort +" user="+ dbUserName +" dbname="+ dbName +" sslmode=disable password="+ dbPass +" host="+ dbHost)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

}


package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Connection struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func ConnectToDatabase(conn Connection) (*sqlx.DB, error) {
	//connect to a PostgreSQL database
	// Replace the connection details (user, dbname, password, host, Port) with your own
	dataSourceName := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%d sslmode=disable",
		conn.User,
		conn.DBName,
		conn.Password,
		conn.Host,
		conn.Port)

	db, err := sqlx.Connect("postgres", dataSourceName)

	if err != nil {
		log.Fatalln("Error When Connecting To database : " + err.Error())
		return nil, err
	}



	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	return db, nil
}

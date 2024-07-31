package main

import (
	"log"
	"pos-acen/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Println("Error On Load Config Error : " + err.Error())
		return
	}

	_, err = config.ConnectToDatabase(config.Connection{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})

	if err != nil {
		log.Println("Error On Connect To Database Err : " + err.Error())
		return
	}
}

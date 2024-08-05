package main

import (
	"log"
	"pos-acen/internal/modules/users/handler/rest"
	userRepo "pos-acen/internal/modules/users/repository"
	userService "pos-acen/internal/modules/users/service"
	"pos-acen/internal/routes"
	"pos-acen/pkg/config"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/thedevsaddam/renderer"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Println("Error On Load Config Error : " + err.Error())
		return
	}

	mysql, err := config.ConnectToDatabase(config.Connection{
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

	defer mysql.Close()

	mutex := &sync.Mutex{}
	validator := validator.New()
	render := renderer.New()
	routes := setupRoutes(render, mysql, validator, cfg, mutex)
	routes.Run(cfg.AppPort)
}

func setupRoutes(render *renderer.Render, myDb *sqlx.DB, validator *validator.Validate, config *config.Config, mutex *sync.Mutex) *routes.Routes {
	userRepo := userRepo.NewUserRepository(myDb)
	userService := userService.NewUserService(userRepo)
	userHandler := rest.NewUserHandler(userService, render)

	return &routes.Routes{
		User: userHandler,
	}
}

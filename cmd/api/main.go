package main

import (
	"Expense-Tracker-API/config/cfg"
	"Expense-Tracker-API/internal/controller"
	"Expense-Tracker-API/internal/repository/postgres"
	"Expense-Tracker-API/internal/router"
	"Expense-Tracker-API/internal/service"
	"fmt"
	"log"

	"github.com/labstack/echo"
)

func main() {
	config, err := cfg.LoadConfig("config/cfg/config.yaml")
	if err != nil {
		log.Fatal("Error load config: " + err.Error())
	}
	log.Println(config)

	repo := postgres.NewPostgresRepo(fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=disable",
		config.DataBase.User, config.DataBase.Password, config.DataBase.DBName, config.DataBase.Port))

	service := service.NewServiceManager(repo)

	controller := controller.NewControllerManager(service)

	e := echo.New()
	router.InitRoutes(e, controller)

	e.Start(config.Server.Port)
}

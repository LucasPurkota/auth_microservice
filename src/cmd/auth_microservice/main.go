package main

import (
	"github.com/LucasPurkota/auth_microservice/internal/config"
	"github.com/LucasPurkota/auth_microservice/internal/database"
	"github.com/LucasPurkota/auth_microservice/routes"
	"fmt"
)

func main() {
	config.LoadConfig()
	database.GORMConnect()

	fmt.Println("Starting in port %s", config.Env.Port)

	r := routes.SetupRoutes()
	r.Run(":" + config.Env.Port)
}

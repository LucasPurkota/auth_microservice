package main

import (
	"auth_microservice/internal/config"
	"auth_microservice/internal/database"
	"auth_microservice/routes"
	"fmt"
)

func main() {
	config.LoadConfig()
	database.GORMConnect()

	fmt.Println("Starting in port %s", config.Env.Port)

	r := routes.SetupRoutes()
	r.Run(":" + config.Env.Port)
}

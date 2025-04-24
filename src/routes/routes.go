package routes

import (
	"time"

	"github.com/LucasPurkota/auth_microservice/internal/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}

func SetupRoutes() *gin.Engine {
	route := gin.Default()

	route.Use(cors.New(corsConfig()))

	personFinance := route.Group("/auth_microservice")
	{
		personFinance.GET("login/:email/:senha", controller.Login)
	}

	return route
}

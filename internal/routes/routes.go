package routes

import (
	"time"

	"github.com/LucasPurkota/auth_microservice/internal/controller"
	"github.com/LucasPurkota/auth_microservice/internal/database"
	"github.com/LucasPurkota/auth_microservice/internal/repository"
	"github.com/LucasPurkota/auth_microservice/internal/service"

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
	userRepo := repository.NewUserRepository(database.Gorm)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(userService)
	route := gin.Default()

	route.Use(cors.New(corsConfig()))

	personFinance := route.Group("/auth_microservice")
	{
		personFinance.GET("/health", controller.Health)
		personFinance.POST("/login", authController.Login)
		personFinance.GET("/auth", authController.Auth)
		user := personFinance.Group("/users")
		{
			user.POST("/", userController.CreateUser)
			user.PUT("/:id", userController.UpdateUser)
			user.PATCH("/:id/password", userController.UpdatePassword)
			user.DELETE("/:id", userController.DeleteUser)
		}
	}

	return route
}

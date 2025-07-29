package routes

import (
	"net/http"
	"time"

	"github.com/LucasPurkota/auth_microservice/internal/controller"
	"github.com/LucasPurkota/auth_microservice/internal/util"

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
		personFinance.GET("/login/:email/:senha", controller.Login)
		user := personFinance.Group("/users")
		{
			user.POST("/", controller.CreateUser)
			user.PUT("/", auth, controller.UpdateUser)
			user.DELETE("/:id", auth, controller.DeleteUser)
		}
	}

	return route
}

func auth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		c.Abort()
		return
	}

	_, err := util.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.Next()
}

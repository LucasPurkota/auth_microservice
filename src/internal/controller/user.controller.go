package controller

import (
	"net/http"

	"github.com/LucasPurkota/auth_microservice/internal/adapter"
	"github.com/LucasPurkota/auth_microservice/internal/database"
	"github.com/LucasPurkota/auth_microservice/internal/model/dto"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userCreated dto.UserCreated

	if err := c.ShouldBindJSON(&userCreated); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Json"})
		return
	}

	user := adapter.UserCreatedToEntity(userCreated)

	query := database.Gorm.Create(&user)
	if query.Error != nil {
		c.JSON(500, gin.H{"error": "Create user error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "User created successfully"})
}

func UpdateUser(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{"data": "Ok"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id is required"})
		return
	}

	query := database.Gorm.Delete(&dto.UserResponse{}, id)
	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": query.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "Ok"})
}

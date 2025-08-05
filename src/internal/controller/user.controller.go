package controller

import (
	"net/http"

	"github.com/LucasPurkota/auth_microservice/internal/adapter"
	"github.com/LucasPurkota/auth_microservice/internal/database"
	"github.com/LucasPurkota/auth_microservice/internal/model/dto"
	"github.com/LucasPurkota/auth_microservice/internal/model/entity"
	"github.com/LucasPurkota/auth_microservice/internal/util"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userCreated dto.UserCreated

	if err := c.ShouldBindJSON(&userCreated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json"})
		return
	}

	user := adapter.UserCreatedToEntity(userCreated)

	selectUser := database.Gorm.Table("public.user").Where("email = ?", user.Email).First(&entity.User{})
	if selectUser.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	query := database.Gorm.Create(&user)
	if query.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create user error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "User created successfully"})
}

func UpdateUser(c *gin.Context) {
	_, err := util.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id is required"})
		return
	}

	var userUpdate dto.UserUpdate
	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json"})
		return
	}

	user := adapter.UserUpdateToEntity(userUpdate)

	query := database.Gorm.Table("public.user").Where("user_id = ?", id).Updates(&user)
	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": query.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Ok"})
}

func UpdatePassword(c *gin.Context) {
	_, err := util.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id is required"})
		return
	}

	var user dto.UserUpdatePassword
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json"})
		return
	}

	if user.NewPassword == user.CurrentPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password cannot be the same as the current password"})
		return
	}

	password, err := util.EncriptedPassword(user.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := database.Gorm.Table("public.user").Where("user_id = ?", id).Update("password", password)
	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": query.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Ok"})
}

func DeleteUser(c *gin.Context) {
	_, err := util.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id is required"})
		return
	}

	query := database.Gorm.Where("user_id = ?", id).Delete(&entity.User{})
	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": query.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Ok"})
}

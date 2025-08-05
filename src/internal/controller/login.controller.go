package controller

import (
	"net/http"

	"github.com/LucasPurkota/auth_microservice/internal/database"
	"github.com/LucasPurkota/auth_microservice/internal/model/dto"
	"github.com/LucasPurkota/auth_microservice/internal/model/entity"
	"github.com/LucasPurkota/auth_microservice/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var userLogin dto.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json"})
		return
	}

	var user entity.User

	query := database.Gorm.Table("public.user").
		Where("email = ?", userLogin.Email).
		First(&user)

	if query.Error != nil && query.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching user"})
		return
	} else if query.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "email or password is incorrect"})
		return
	}

	if !util.VerifyPassword(userLogin.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email or password is incorrect"})
		return
	}

	token, err := util.GenerateJWT(user.UserId, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": token,
	})
}

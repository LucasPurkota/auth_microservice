package controller

import (
	"net/http"

	"github.com/LucasPurkota/auth_microservice/internal/model"
	"github.com/LucasPurkota/auth_microservice/internal/service"
	"github.com/LucasPurkota/auth_microservice/internal/util"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.UserService
}

func NewAuthController(authService *service.UserService) *AuthController {
	return &AuthController{authService: authService}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var credentials model.UserLogin
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json"})
		return
	}

	token, err := ctrl.authService.Login(c.Request.Context(), credentials)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "email or password is incorrect" {
			statusCode = http.StatusUnauthorized
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": token,
	})
}

func (ctrl *AuthController) Auth(c *gin.Context) {
	_, err := util.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Authorized"})
}

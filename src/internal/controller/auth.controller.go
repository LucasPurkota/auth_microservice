package controller

import (
	"net/http"

	"github.com/LucasPurkota/auth_microservice/internal/util"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	_, err := util.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Authorized"})
}

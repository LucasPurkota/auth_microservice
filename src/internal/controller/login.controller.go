package controller

import (
	"github.com/LucasPurkota/auth_microservice/internal/database"
	"github.com/LucasPurkota/auth_microservice/internal/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func gerarHashSenha(senha string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return string(hash), err
}

func verificarSenha(senha, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	return err == nil
}

func Login(c *gin.Context) {
	email := c.Param("email")

	senhas := c.Param("senha")

	var senha string
	query := database.Gorm.Table("public.usuario").Model(model.Usuario{}).Select("senha").Where("email = ?", email).First(&senha)
	if err := query.Error; err != nil {
		c.JSON(500, gin.H{"error": "Erro ao consultar o banco de dados"})
		return
	} else if query.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Usuário não encontrado"})
		return
	}

	if !verificarSenha(senhas, senha) {
		c.JSON(401, gin.H{"error": "Autenticação inválida"})
		return
	}
	c.JSON(200, "autenticado")
}

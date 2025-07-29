package controller

import (
	"context"

	"github.com/LucasPurkota/auth_microservice/internal/database"
	"github.com/LucasPurkota/auth_microservice/internal/model/entity"
	"github.com/LucasPurkota/auth_microservice/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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

var ctx = context.Background()

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func Login(c *gin.Context) {
	email := c.Param("email")
	senhas := c.Param("senha")

	// cacheKey := "usuario:senha:" + email
	// val, err := redisClient.Get(ctx, cacheKey).Result()

	var usuario entity.User

	// if err == redis.Nil {
	query := database.Gorm.Table("public.user").
		Model(entity.User{}).
		Where("email = ?", email).
		First(&usuario)

	if query.Error != nil {
		c.JSON(500, gin.H{"error": "Erro ao consultar o banco de dados"})
		return  
	} else if query.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Usuário não encontrado"})
		return
	}

	// 	userJSON, err := json.Marshal(usuario)
	// 	if err != nil {
	// 		c.JSON(500, gin.H{"error": "Erro ao serializar usuário"})
	// 		return
	// 	}

	// 	err = redisClient.Set(ctx, cacheKey, userJSON, 1*time.Minute).Err()
	// 	if err != nil {
	// 		c.JSON(500, gin.H{"error": "Erro ao salvar no cache"})
	// 		return
	// 	}
	// } else if err != nil {
	// 	c.JSON(500, gin.H{"error": "Erro no cache: " + err.Error()})
	// 	return
	// } else {
	// 	err = json.Unmarshal([]byte(val), &usuario)
	// 	if err != nil {
	// 		c.JSON(500, gin.H{"error": "Erro ao desserializar usuário do cache"})
	// 		return
	// 	}
	// }

	if !verificarSenha(senhas, usuario.Password) {
		c.JSON(401, gin.H{"autentication": "Autenticação inválida"})
		return
	}

	token, err := util.GenerateJWT(usuario.UserId, usuario.Email, usuario.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(200, gin.H{
		"data":  usuario,
		"Token": token,
	})
}

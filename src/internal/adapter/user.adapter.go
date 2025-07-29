package adapter

import (
	"database/sql"
	"time"

	"github.com/LucasPurkota/auth_microservice/internal/model/dto"
	"github.com/LucasPurkota/auth_microservice/internal/model/entity"
	"github.com/LucasPurkota/auth_microservice/internal/util"
)

func UserCreatedToEntity(userCreated dto.UserCreated) entity.User {
	encriptPassword, err := util.EncriptedPassword(userCreated.Password)
	if err != nil {
		return entity.User{}
	}
	return entity.User{
		Name:      userCreated.Name,
		LastName:  userCreated.LastName,
		Email:     userCreated.Email,
		Password:  encriptPassword,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}
}

func UserEntityToResponse(user entity.User) dto.UserResponse {
	return dto.UserResponse{
		Name:     user.Name,
		LastName: user.LastName,
	}
}

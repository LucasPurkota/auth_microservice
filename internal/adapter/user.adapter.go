package adapter

import (
	"database/sql"
	"time"

	"github.com/LucasPurkota/auth_microservice/internal/model"
	"github.com/LucasPurkota/auth_microservice/internal/util"
)

func UserCreatedToEntity(userCreated model.UserCreated) model.User {
	encriptPassword, err := util.EncriptedPassword(userCreated.Password)
	if err != nil {
		return model.User{}
	}
	return model.User{
		Name:      userCreated.Name,
		LastName:  userCreated.LastName,
		Email:     userCreated.Email,
		Password:  encriptPassword,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}
}

func UserUpdateToEntity(userUpdate model.UserUpdate) model.User {
	return model.User{
		Name:     userUpdate.Name,
		LastName: userUpdate.LastName,
		Email:    userUpdate.Email,
	}
}

func UserEntityToResponse(user model.User) model.UserResponse {
	return model.UserResponse{
		Name:     user.Name,
		LastName: user.LastName,
	}
}

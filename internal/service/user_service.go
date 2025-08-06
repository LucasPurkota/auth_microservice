package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/LucasPurkota/auth_microservice/internal/model"
	"github.com/LucasPurkota/auth_microservice/internal/util"
)

type UserService struct {
	repo model.UserRepository
}

func NewUserService(repo model.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, input model.UserCreated) error {
	// Verifica se email já existe
	existing, err := s.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("email already registered")
	}

	// Criptografa senha
	hashedPassword, err := util.EncriptedPassword(input.Password)
	if err != nil {
		return err
	}

	user := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	return s.repo.Create(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, id string, input model.UserUpdate) error {
	user := &model.User{
		Name:  input.Name,
		Email: input.Email,
	}
	return s.repo.Update(ctx, id, user)
}

func (s *UserService) UpdatePassword(ctx context.Context, id string, input model.UserUpdatePassword) error {
	if input.NewPassword == input.CurrentPassword {
		return errors.New("new password cannot be the same as the current password")
	}

	hashedPassword, err := util.EncriptedPassword(input.NewPassword)
	if err != nil {
		return err
	}

	return s.repo.UpdatePassword(ctx, id, hashedPassword)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) Login(ctx context.Context, credentials model.UserLogin) (string, error) {
	// Busca usuário com senha
	user, err := s.repo.GetByEmailWithPassword(ctx, credentials.Email)
	if err != nil {
		return "", fmt.Errorf("error fetching user: %w", err)
	}
	if user == nil {
		return "", errors.New("email or password is incorrect")
	}

	// Verifica senha
	if !util.VerifyPassword(credentials.Password, user.Password) {
		return "", errors.New("email or password is incorrect")
	}

	return util.GenerateJWT(user.UserId, user.Email, user.Password)
}

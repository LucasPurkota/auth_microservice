package repository

import (
	"context"
	"errors"

	"github.com/LucasPurkota/auth_microservice/internal/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	result := r.db.Table("public.user").Create(user)
	return result.Error
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	result := r.db.Table("public.user").Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, result.Error
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	result := r.db.Table("public.user").Where("user_id = ?", id).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, result.Error
}

func (r *userRepository) Update(ctx context.Context, id string, user *model.User) error {
	result := r.db.Table("public.user").Where("user_id = ?", id).Updates(user)
	return result.Error
}

func (r *userRepository) UpdatePassword(ctx context.Context, id, newPassword string) error {
	result := r.db.Table("public.user").Where("user_id = ?", id).Update("password", newPassword)
	return result.Error
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	result := r.db.Table("public.user").Where("user_id = ?", id).Delete(&model.User{})
	return result.Error
}

func (r *userRepository) GetByEmailWithPassword(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	result := r.db.Table("public.user").
		Select("user_id, email, password").
		Where("email = ?", email).
		First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, result.Error
}

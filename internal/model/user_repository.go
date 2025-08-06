package model

import "context"

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, id string, user *User) error
	UpdatePassword(ctx context.Context, id, newPassword string) error
	Delete(ctx context.Context, id string) error
	GetByEmailWithPassword(ctx context.Context, email string) (*User, error)
}

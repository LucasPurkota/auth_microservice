package model

import "database/sql"

type User struct {
	UserId    string `gorm:"primaryKey;default:gen_random_uuid()"`
	Name      string
	LastName  string
	Email     string
	Password  string
	CreatedAt sql.NullTime
}

func (User) TableName() string {
	return "public.user"	
}

type UserCreated struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

type UserUpdatePassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type UserResponse struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

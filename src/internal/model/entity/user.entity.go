package entity

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

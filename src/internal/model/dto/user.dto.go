package dto

type UserCreated struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

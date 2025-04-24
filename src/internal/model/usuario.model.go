package model

type Usuario struct {
	UsuarioId int    `json:"usuario_id"`
	Nome      string `json:"nome"`
	Senha     string `json:"senha"`
	Email     string `json:"email"`
}

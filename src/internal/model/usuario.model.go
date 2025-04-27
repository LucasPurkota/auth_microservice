package model

type Usuario struct {
	UsuarioId string `json:"usuario_id"`
	Nome      string `json:"nome"`
	Senha     string `json:"senha"`
	Email     string `json:"email"`
}

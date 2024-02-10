package auth

import "time"

type LoginRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Name     string `validate:"required,max=100" json:"name"`
	Email    string `validate:"required,min=8,max=100" json:"email"`
	Password string `validate:"required,min=6,max=100" json:"password"`
}

type RegisterResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

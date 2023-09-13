package auth

type RegisterRequest struct {
	Name     string `validate:"required,max=100" json:"name"`
	Email    string `validate:"required,min=8,max=100" json:"email"`
	Password string `validate:"required,min=6,max=100" json:"password"`
}

package user

type CreateRequest struct {
	Name     string `validate:"required,max=100,min=1" json:"name"`
	Email    string `validate:"required,max=100,min=1" json:"email"`
	Password string `validate:"required,max=225,min=8" json:"password"`
}

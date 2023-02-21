package api

type AdminCreateRequest struct {
	Name     string `validate:"required,min=3,max=20"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

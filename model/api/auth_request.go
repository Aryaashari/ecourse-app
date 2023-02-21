package api

type AuthRegisterRequest struct {
	Name     string `validate:"required,min=3,max=20"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type AuthLoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

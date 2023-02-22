package api

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdateRequest struct {
	Id   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

package service

import (
	"context"
	"database/sql"
	"ecourse-app/model/api"
	"ecourse-app/repository"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	FindAll(ctx context.Context) []api.UserResponse
	FindById(ctx context.Context, id int64) api.ApiResponse
	Create(ctx context.Context, request api.UserCreateRequest) api.ApiResponse
	Update(ctx context.Context, request api.UserUpdateRequest) api.ApiResponse
	Delete(ctx context.Context, id int64)
}

type UserServiceImpl struct {
	UserRepo  repository.UserRepository
	DB        *sql.DB
	Validator *validator.Validate
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []api.UserResponse
func (service *UserServiceImpl) FindById(ctx context.Context, id int64) api.ApiResponse
func (service *UserServiceImpl) Create(ctx context.Context, request api.UserCreateRequest) api.ApiResponse
func (service *UserServiceImpl) Update(ctx context.Context, request api.UserUpdateRequest) api.ApiResponse
func (service *UserServiceImpl) Delete(ctx context.Context, id int64)

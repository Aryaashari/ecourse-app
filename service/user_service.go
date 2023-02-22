package service

import (
	"context"
	"database/sql"
	"ecourse-app/exception"
	"ecourse-app/helper"
	"ecourse-app/model/api"
	"ecourse-app/model/domain"
	"ecourse-app/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	FindAll(ctx context.Context) []api.UserResponse
	FindById(ctx context.Context, id int64) api.UserResponse
	Create(ctx context.Context, request api.UserCreateRequest) api.UserResponse
	Update(ctx context.Context, request api.UserUpdateRequest) api.UserResponse
	Delete(ctx context.Context, id int64)
}

type UserServiceImpl struct {
	UserRepo repository.UserRepository
	DB       *sql.DB
	Validate *validator.Validate
}

func NewUserService(userRepo repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
		DB:       db,
		Validate: validate,
	}
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []api.UserResponse {
	transaction, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(transaction)

	users := service.UserRepo.FindAll(ctx, transaction)

	var userResponses []api.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, helper.ConvertToUserResponse(&user))
	}

	return userResponses
}

func (service *UserServiceImpl) FindById(ctx context.Context, id int64) api.UserResponse {
	transaction, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(transaction)

	user, err := service.UserRepo.FindById(ctx, transaction, id)
	helper.PanicError(err)

	return helper.ConvertToUserResponse(&user)
}

func (service *UserServiceImpl) Create(ctx context.Context, request api.UserCreateRequest) api.UserResponse {
	// Validation
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	transaction, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(transaction)

	// Check existing email, if email was exist then return error
	_, err = service.UserRepo.FindByEmail(ctx, transaction, request.Email)
	if err == nil {
		panic(exception.NewConflictError("this email is already registered"))
	}

	// Encrypt Password
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	helper.PanicError(err)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(bytes),
	}

	user = service.UserRepo.Insert(ctx, transaction, user)
	return helper.ConvertToUserResponse(&user)

}

func (service *UserServiceImpl) Update(ctx context.Context, request api.UserUpdateRequest) api.UserResponse {
	// Validation
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	transaction, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(transaction)

	user, err := service.UserRepo.FindById(ctx, transaction, request.Id)
	helper.PanicError(err)

	user.Name = request.Name

	user = service.UserRepo.Update(ctx, transaction, user)
	return helper.ConvertToUserResponse(&user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, id int64) {
	transaction, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(transaction)

	_, err = service.UserRepo.FindById(ctx, transaction, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepo.Delete(ctx, transaction, id)
}

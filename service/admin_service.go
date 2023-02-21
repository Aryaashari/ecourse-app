package service

import (
	"context"
	"database/sql"
	"ecourse-app/helper"
	"ecourse-app/model/api"
	"ecourse-app/model/domain"
	"ecourse-app/repository"
	"errors"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	Create(ctx context.Context, request api.AdminCreateRequest) (api.AdminResponse, error)
	FindByEmail(ctx context.Context, email string) api.AdminResponse
}

type AdminServiceImpl struct {
	AdminRepo repository.AdminRepository
	DB        *sql.DB
	Validate  *validator.Validate
}

func NewAdminService(adminRepo repository.AdminRepository, db *sql.DB, validate *validator.Validate) AdminService {
	return &AdminServiceImpl{
		AdminRepo: adminRepo,
		DB:        db,
		Validate:  validate,
	}
}

func (service *AdminServiceImpl) Create(ctx context.Context, request api.AdminCreateRequest) (api.AdminResponse, error) {
	// Validation
	err := service.Validate.Struct(request)
	helper.PanicError(&err)

	transaction, err := service.DB.Begin()
	helper.PanicError(&err)
	defer helper.CommitRollback(transaction)

	// Check existing email, if email was exist then return error
	_, err = service.AdminRepo.FindByEmail(ctx, transaction, request.Email)
	if err == nil {
		return api.AdminResponse{}, errors.New("email was registered")
	}

	// Encrypt Password
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 20)
	helper.PanicError(&err)

	admin := domain.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(bytes),
	}

	admin = service.AdminRepo.Insert(ctx, transaction, admin)
	return helper.ConvertToAdminResponse(&admin), nil
}

func (service *AdminServiceImpl) FindByEmail(ctx context.Context, email string) api.AdminResponse {

	transaction, err := service.DB.Begin()
	helper.PanicError(&err)

	defer helper.CommitRollback(transaction)

	admin, err := service.AdminRepo.FindByEmail(ctx, transaction, email)
	helper.PanicError(&err)

	return helper.ConvertToAdminResponse(&admin)
}

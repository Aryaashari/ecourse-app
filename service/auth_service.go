package service

import (
	"context"
	"database/sql"
	"ecourse-app/config"
	"ecourse-app/exception"
	"ecourse-app/helper"
	"ecourse-app/model/api"
	"ecourse-app/model/domain"
	"ecourse-app/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, request api.AuthRegisterRequest) api.AuthResponse
	// FindByEmail(ctx context.Context, email string) api.AdminResponse
	Login(ctx context.Context, request api.AuthLoginRequest) string // return token
}

type AuthServiceImpl struct {
	AdminRepo repository.AdminRepository
	DB        *sql.DB
	Validate  *validator.Validate
}

func NewAuthService(adminRepo repository.AdminRepository, db *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AdminRepo: adminRepo,
		DB:        db,
		Validate:  validate,
	}
}

func (service *AuthServiceImpl) Register(ctx context.Context, request api.AuthRegisterRequest) api.AuthResponse {
	// Validation
	err := service.Validate.Struct(request)
	helper.PanicError(&err)

	transaction, err := service.DB.Begin()
	helper.PanicError(&err)
	defer helper.CommitRollback(transaction)

	// Check existing email, if email was exist then return error
	_, err = service.AdminRepo.FindByEmail(ctx, transaction, request.Email)
	if err == nil {
		panic(exception.NewConflictError("this email is already registered"))
	}

	// Encrypt Password
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	helper.PanicError(&err)

	admin := domain.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(bytes),
	}

	admin = service.AdminRepo.Insert(ctx, transaction, admin)
	return helper.ConvertToAuthResponse(&admin)
}

func (service *AuthServiceImpl) Login(ctx context.Context, request api.AuthLoginRequest) string {

	// Validation
	err := service.Validate.Struct(request)
	helper.PanicError(&err)

	transaction, err := service.DB.Begin()
	helper.PanicError(&err)
	defer helper.CommitRollback(transaction)

	// Check existing email, if email not registered then return error
	admin, err := service.AdminRepo.FindByEmail(ctx, transaction, request.Email)
	if err != nil {
		panic(exception.NewUnauthorizedError("email or password invalid"))
	}

	// Check password, if password is not match then return error
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewUnauthorizedError("email or password invalid"))
	}

	// Create Claims
	claims := config.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "E-Course App",
			ExpiresAt: time.Now().Add(config.JWT_EXPIRED).Unix(),
		},
		Name:  admin.Name,
		Email: admin.Email,
	}

	// Generate Token
	token := jwt.NewWithClaims(config.JWT_SIGNING_METHOD, claims)
	signedToken, err := token.SignedString(config.JWT_KEY)
	helper.PanicError(&err)

	return signedToken
}

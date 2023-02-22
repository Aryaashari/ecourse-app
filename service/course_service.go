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
)

type CourseService interface {
	FindAll(ctx context.Context) []api.CourseResponse
	FindById(ctx context.Context, id int64) api.CourseResponse
	Create(ctx context.Context, request api.CourseCreateRequest) api.CourseResponse
	Update(ctx context.Context, request api.CourseUpdateRequest) api.CourseResponse
	Delete(ctx context.Context, id int64)
}

type CourseServiceImpl struct {
	CourseRepo         repository.CourseRepository
	CourseCategoryRepo repository.CourseCategoryRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

func NewCourseService(courseRepo repository.CourseRepository, courseCategoryRepo repository.CourseCategoryRepository, db *sql.DB, validator *validator.Validate) CourseService {
	return &CourseServiceImpl{
		CourseRepo:         courseRepo,
		CourseCategoryRepo: courseCategoryRepo,
		DB:                 db,
		Validator:          validator,
	}
}

func (service *CourseServiceImpl) FindAll(ctx context.Context) []api.CourseResponse {

	transaction, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(transaction)

	var courseResponses []api.CourseResponse
	courses, courseCategories := service.CourseRepo.FindAll(ctx, transaction)
	for index, course := range courses {
		courseResponses = append(courseResponses, helper.ConvertToCourseResponse(&course, &domain.CourseCategory{
			Id:   courseCategories[index].Id,
			Name: courseCategories[index].Name,
		}))
	}

	return courseResponses

}

func (service *CourseServiceImpl) FindById(ctx context.Context, id int64) api.CourseResponse {
	transaction, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(transaction)

	course, err := service.CourseRepo.FindById(ctx, transaction, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	courseCategory, err := service.CourseCategoryRepo.FindById(ctx, transaction, course.CourseCategoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Get all users by course id
	// ...

	return helper.ConvertToCourseResponse(&course, &courseCategory)
}

func (service *CourseServiceImpl) Create(ctx context.Context, request api.CourseCreateRequest) api.CourseResponse {
	// Validation
	err := service.Validator.Struct(request)
	helper.PanicError(err)

	transaction, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(transaction)

	courseCategory, err := service.CourseCategoryRepo.FindById(ctx, transaction, request.CourseCategoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	course := domain.Course{
		Title:            request.Title,
		CourseCategoryId: request.CourseCategoryId,
	}
	course = service.CourseRepo.Insert(ctx, transaction, course)

	return helper.ConvertToCourseResponse(&course, &courseCategory)

}

func (service *CourseServiceImpl) Update(ctx context.Context, request api.CourseUpdateRequest) api.CourseResponse {
	// Validation
	err := service.Validator.Struct(request)
	helper.PanicError(err)

	transaction, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(transaction)

	course, err := service.CourseRepo.FindById(ctx, transaction, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	courseCategory, err := service.CourseCategoryRepo.FindById(ctx, transaction, request.CourseCategoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	course.Title = request.Title
	course.CourseCategoryId = request.CourseCategoryId

	course = service.CourseRepo.Update(ctx, transaction, course)

	return helper.ConvertToCourseResponse(&course, &courseCategory)
}

func (service *CourseServiceImpl) Delete(ctx context.Context, id int64) {
	transaction, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitRollback(transaction)

	_, err = service.CourseRepo.FindById(ctx, transaction, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CourseRepo.Delete(ctx, transaction, id)
}

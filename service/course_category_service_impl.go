package service

import (
	"context"
	"database/sql"
	"ecourse-app/helper"
	api "ecourse-app/model/api/course_category"
	"ecourse-app/model/domain"
	"ecourse-app/repository"

	"github.com/go-playground/validator/v10"
)

type CourseCategoryServiceImpl struct {
	CourseCategoryRepo repository.CourseCategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func (service *CourseCategoryServiceImpl) FindAll(ctx context.Context) []api.CourseCategoryResponse {
	transaction, err := service.DB.Begin()
	helper.PanicError(&err)

	defer helper.CommitRollback(transaction)

	courseCategories := service.CourseCategoryRepo.FindAll(ctx, transaction)

	var courseCategoryResponses []api.CourseCategoryResponse
	for _, courseCategory := range courseCategories {
		courseCategoryResponses = append(courseCategoryResponses, helper.ConvertToCourseCategoryResponse(&courseCategory))
	}

	return courseCategoryResponses
}

// func (service *CourseCategoryServiceImpl) FindById(ctx context.Context, id int64) api.CourseCategoryResponse {
// 	transaction, err := service.DB.Begin()
// 	helper.PanicError(&err)

// 	defer helper.CommitRollback(transaction)

// 	courseCategory, err := service.CourseCategoryRepo.FindById(ctx, transaction, id)
// 	helper.PanicError(&err)

// }

func (service *CourseCategoryServiceImpl) Create(ctx context.Context, request api.CourseCategoryCreateRequest) api.CourseCategoryResponse {
	// Validation
	err := service.Validate.Struct(request)
	helper.PanicError(&err)

	transaction, err := service.DB.Begin()
	helper.PanicError(&err)

	defer helper.CommitRollback(transaction)

	courseCategory := domain.CourseCategory{
		Name: request.Name,
	}

	courseCategory = service.CourseCategoryRepo.Insert(ctx, transaction, courseCategory)

	return helper.ConvertToCourseCategoryResponse(&courseCategory)
}

func (service *CourseCategoryServiceImpl) Update(ctx context.Context, request api.CourseCategoryUpdateRequest) api.CourseCategoryResponse {
	// Validation
	err := service.Validate.Struct(request)
	helper.PanicError(&err)

	transaction, err := service.DB.Begin()
	helper.PanicError(&err)

	defer helper.CommitRollback(transaction)

	courseCategory, err := service.CourseCategoryRepo.FindById(ctx, transaction, request.Id)
	helper.PanicError(&err)

	courseCategory.Name = request.Name

	courseCategory = service.CourseCategoryRepo.Update(ctx, transaction, courseCategory)

	return helper.ConvertToCourseCategoryResponse(&courseCategory)
}

func (service *CourseCategoryServiceImpl) Delete(ctx context.Context, id int64) {
	transaction, err := service.DB.Begin()
	helper.PanicError(&err)

	defer helper.CommitRollback(transaction)

	courseCategory, err := service.CourseCategoryRepo.FindById(ctx, transaction, id)
	helper.PanicError(&err)

	service.CourseCategoryRepo.Delete(ctx, transaction, courseCategory.Id)
}

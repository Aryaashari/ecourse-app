package service

import (
	"context"
	api "ecourse-app/model/api/course_category"
)

type CourseCategoryService interface {
	FindAll(ctx context.Context) []api.CourseCategoryResponse
	FindById(ctx context.Context, id int64) api.CourseCategoryResponse // (api.CourseCategoryResponse, api.CourseResponse)
	Create(ctx context.Context, request api.CourseCategoryCreateRequest) api.CourseCategoryResponse
	Update(ctx context.Context, request api.CourseCategoryUpdateRequest) api.CourseCategoryResponse
	Delete(ctx context.Context, id int64)
}

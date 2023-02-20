package helper

import (
	api "ecourse-app/model/api/course_category"
	"ecourse-app/model/domain"
)

func ConvertToCourseCategoryResponse(courseCategory *domain.CourseCategory) api.CourseCategoryResponse {
	return api.CourseCategoryResponse{
		Id:   courseCategory.Id,
		Name: courseCategory.Name,
	}
}

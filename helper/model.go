package helper

import (
	"ecourse-app/model/api"
	"ecourse-app/model/domain"
)

func ConvertToCourseCategoryResponse(courseCategory *domain.CourseCategory) api.CourseCategoryResponse {
	return api.CourseCategoryResponse{
		Id:   courseCategory.Id,
		Name: courseCategory.Name,
	}
}

func ApiResponseFormatter(status int, message string, data interface{}) api.ApiResponse {
	return api.ApiResponse{
		Info: api.InfoField{
			Status:  status,
			Message: message,
		},
		Data: data,
	}
}

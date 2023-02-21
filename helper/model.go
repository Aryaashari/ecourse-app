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

func ApiResponseFormatter(code int, status string, message string, data interface{}) api.ApiResponse {
	return api.ApiResponse{
		Info: api.InfoField{
			Code:    code,
			Status:  status,
			Message: message,
		},
		Data: data,
	}
}

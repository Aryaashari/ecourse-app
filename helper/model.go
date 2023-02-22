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

func ConvertToAuthResponse(admin *domain.Admin) api.AuthResponse {
	return api.AuthResponse{
		Id:    admin.Id,
		Name:  admin.Name,
		Email: admin.Email,
	}
}

func ConvertToCourseResponse(course *domain.Course, courseCategory *domain.CourseCategory) api.CourseResponse {
	return api.CourseResponse{
		Id:       course.Id,
		Title:    course.Title,
		Category: ConvertToCourseCategoryResponse(courseCategory),
	}
}

func ConvertToUserResponse(user *domain.User) api.UserResponse {
	return api.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
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

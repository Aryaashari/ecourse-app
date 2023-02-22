package api

type CourseCreateRequest struct {
	CourseCategoryId int64  `json:"course_category_id" validate:"required"`
	Title            string `json:"title" validate:"required,min=3,max=50"`
}

type CourseUpdateRequest struct {
	Id               int64  `validate:"required"`
	CourseCategoryId int64  `json:"course_category_id"`
	Title            string `json:"title" validate:"required,min=3,max=50"`
}

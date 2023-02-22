package api

type CourseResponse struct {
	Id       int64                  `json:"id"`
	Title    string                 `json:"title"`
	Category CourseCategoryResponse `json:"category"`
}

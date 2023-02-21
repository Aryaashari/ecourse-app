package api

type CourseCategoryCreateRequest struct {
	Name string `validate:"required,min=3,max=30" json:"name"`
}

type CourseCategoryUpdateRequest struct {
	Id   int64  `validate:"required" json:"id"`
	Name string `validate:"required,min=3,max=30" json:"name"`
}

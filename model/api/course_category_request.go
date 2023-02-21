package api

type CourseCategoryCreateRequest struct {
	Name string `validate:"required, min=3, max=30"`
}

type CourseCategoryUpdateRequest struct {
	Id   int64  `validate:"required"`
	Name string `validate:"required, min=3. max=30"`
}

package api

type CourseCategoryCreateRequest struct {
	Name string
}

type CourseCategoryUpdateRequest struct {
	Id   int64
	Name string
}

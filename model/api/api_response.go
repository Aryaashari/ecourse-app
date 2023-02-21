package api

type InfoField struct {
	Status  int
	Message string
}

type ApiResponse struct {
	Info InfoField
	Data interface{}
}

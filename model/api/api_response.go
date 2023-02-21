package api

type InfoField struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ApiResponse struct {
	Info InfoField   `json:"info"`
	Data interface{} `json:"data"`
}

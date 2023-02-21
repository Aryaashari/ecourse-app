package api

type InfoField struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponse struct {
	Info InfoField   `json:"info"`
	Data interface{} `json:"data"`
}

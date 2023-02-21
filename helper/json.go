package helper

import (
	"encoding/json"
	"net/http"
)

func HandleRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicError(&err)
}

func HandleApiResponse(writer http.ResponseWriter, apiResponse interface{}) {
	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(apiResponse)
	PanicError(&err)
}

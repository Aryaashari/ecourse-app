package exception

import (
	"ecourse-app/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	} else if validationError(writer, request, err) {
		return
	} else if conflictError(writer, request, err) {
		return
	} else if unauthorizedError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)

}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		apiResponse := helper.ApiResponseFormatter(http.StatusUnauthorized, "unauthorized", exception.Error, nil)
		helper.HandleApiResponse(writer, apiResponse)

		return true
	}

	return false
}

func conflictError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(ConflictError)
	if ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusConflict)

		apiResponse := helper.ApiResponseFormatter(http.StatusConflict, "conflict", exception.Error, nil)
		helper.HandleApiResponse(writer, apiResponse)

		return true
	}

	return false
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		apiResponse := helper.ApiResponseFormatter(http.StatusBadRequest, "bad request", exception.Error(), nil)
		helper.HandleApiResponse(writer, apiResponse)

		return true
	}

	return false
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		apiResponse := helper.ApiResponseFormatter(http.StatusNotFound, "not found", exception.Error, nil)
		helper.HandleApiResponse(writer, apiResponse)

		return true
	}

	return false
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(error)

	if ok {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)

		apiResponse := helper.ApiResponseFormatter(http.StatusInternalServerError, "internal server error", exception.Error(), nil)
		helper.HandleApiResponse(writer, apiResponse)
		return true
	}

	return false
}

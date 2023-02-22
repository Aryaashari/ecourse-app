package controller

import (
	"ecourse-app/helper"
	"ecourse-app/model/api"
	"ecourse-app/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userReponses := controller.UserService.FindAll(request.Context())

	apiResponse := helper.ApiResponseFormatter(200, "success", "get all data success", userReponses)

	helper.HandleApiResponse(writer, apiResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("userId"))
	helper.PanicError(err)

	userReponses := controller.UserService.FindById(request.Context(), int64(id))

	apiResponse := helper.ApiResponseFormatter(200, "success", "get detail data success", userReponses)

	helper.HandleApiResponse(writer, apiResponse)
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var userCreateRequest api.UserCreateRequest

	helper.HandleRequestBody(request, &userCreateRequest)

	userReponse := controller.UserService.Create(request.Context(), userCreateRequest)

	apiResponse := helper.ApiResponseFormatter(200, "success", "create user success", userReponse)

	helper.HandleApiResponse(writer, apiResponse)

}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var userUpdateRequest api.UserUpdateRequest

	id, err := strconv.Atoi(params.ByName("userId"))
	helper.PanicError(err)

	helper.HandleRequestBody(request, &userUpdateRequest)

	userUpdateRequest.Id = int64(id)

	userReponse := controller.UserService.Update(request.Context(), userUpdateRequest)

	apiResponse := helper.ApiResponseFormatter(200, "success", "update user success", userReponse)

	helper.HandleApiResponse(writer, apiResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("userId"))
	helper.PanicError(err)

	controller.UserService.Delete(request.Context(), int64(id))

	apiResponse := helper.ApiResponseFormatter(200, "success", "delete data success", nil)

	helper.HandleApiResponse(writer, apiResponse)
}

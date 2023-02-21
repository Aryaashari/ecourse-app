package controller

import (
	"ecourse-app/helper"
	"ecourse-app/model/api"
	"ecourse-app/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var authRegisterRequest api.AuthRegisterRequest
	helper.HandleRequestBody(request, &authRegisterRequest)

	authRegisterReponse := controller.AuthService.Register(request.Context(), authRegisterRequest)

	apiResponse := helper.ApiResponseFormatter(200, "success", "register success", authRegisterReponse)

	helper.HandleApiResponse(writer, apiResponse)

}
func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	var authLoginRequest api.AuthLoginRequest
	helper.HandleRequestBody(request, &authLoginRequest)

	authLoginResponse := controller.AuthService.Login(request.Context(), authLoginRequest)

	apiResponse := helper.ApiResponseFormatter(200, "success", "login success", map[string]string{"token": authLoginResponse})

	helper.HandleApiResponse(writer, apiResponse)
}

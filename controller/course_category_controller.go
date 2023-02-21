package controller

import (
	"ecourse-app/helper"
	"ecourse-app/model/api"
	"ecourse-app/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CourseCategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	// FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type CourseCategoryControllerImpl struct {
	CourseCategoryService service.CourseCategoryService
}

func NewCourseCategoryController(courseCategoryService service.CourseCategoryService) CourseCategoryController {
	return &CourseCategoryControllerImpl{
		CourseCategoryService: courseCategoryService,
	}
}

func (controller *CourseCategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var courseCategoryCreateRequest api.CourseCategoryCreateRequest

	helper.HandleRequestBody(request, &courseCategoryCreateRequest)

	courseCategoryResponse := controller.CourseCategoryService.Create(request.Context(), courseCategoryCreateRequest)

	apiResponse := helper.ApiResponseFormatter(200, "success", "create course category success", courseCategoryResponse)

	helper.HandleApiResponse(writer, apiResponse)
}

func (controller *CourseCategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var courseCategoryUpdateRequest api.CourseCategoryUpdateRequest

	helper.HandleRequestBody(request, &courseCategoryUpdateRequest)

	id, err := strconv.Atoi(params.ByName("courseCategoryId"))
	helper.PanicError(&err)

	courseCategoryUpdateRequest.Id = int64(id)
	courseCategoryResponse := controller.CourseCategoryService.Update(request.Context(), courseCategoryUpdateRequest)

	apiResponse := helper.ApiResponseFormatter(200, "success", "update course category success", courseCategoryResponse)

	helper.HandleApiResponse(writer, apiResponse)
}

func (controller *CourseCategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	id, err := strconv.Atoi(params.ByName("courseCategoryId"))
	helper.PanicError(&err)

	controller.CourseCategoryService.Delete(request.Context(), int64(id))

	apiResponse := helper.ApiResponseFormatter(200, "success", "delete course category success", nil)

	helper.HandleApiResponse(writer, apiResponse)
}

func (controller *CourseCategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	courseCategoryResponses := controller.CourseCategoryService.FindAll(request.Context())

	apiResponse := helper.ApiResponseFormatter(200, "success", "get all course categories success", courseCategoryResponses)

	helper.HandleApiResponse(writer, apiResponse)
}

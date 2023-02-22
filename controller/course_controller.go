package controller

import (
	"ecourse-app/helper"
	"ecourse-app/model/api"
	"ecourse-app/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CourseController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type CourseControllerImpl struct {
	CourseService service.CourseService
}

func NewCourseController(courseService service.CourseService) CourseController {
	return &CourseControllerImpl{
		CourseService: courseService,
	}
}

func (controller *CourseControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	courseReponse := controller.CourseService.FindAll(request.Context())

	apiResponse := helper.ApiResponseFormatter(200, "success", "get all courses success", courseReponse)

	helper.HandleApiResponse(writer, apiResponse)

}

func (controller *CourseControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("courseId"))
	helper.PanicError(err)

	course := controller.CourseService.FindById(request.Context(), int64(id))

	apiResponse := helper.ApiResponseFormatter(200, "success", "get course by id success", course)

	helper.HandleApiResponse(writer, apiResponse)
}

func (controller *CourseControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var courseCreateRequest api.CourseCreateRequest

	helper.HandleRequestBody(request, &courseCreateRequest)

	course := controller.CourseService.Create(request.Context(), courseCreateRequest)

	apiResponse := helper.ApiResponseFormatter(200, "success", "create success", course)

	helper.HandleApiResponse(writer, apiResponse)
}

func (controller *CourseControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var courseUpdateRequest api.CourseUpdateRequest

	helper.HandleRequestBody(request, &courseUpdateRequest)

	id, err := strconv.Atoi(params.ByName("courseId"))
	helper.PanicError(err)

	courseUpdateRequest.Id = int64(id)

	course := controller.CourseService.Update(request.Context(), courseUpdateRequest)

	apiResponse := helper.ApiResponseFormatter(200, "success", "update success", course)

	helper.HandleApiResponse(writer, apiResponse)
}

func (controller *CourseControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	id, err := strconv.Atoi(params.ByName("courseId"))
	helper.PanicError(err)

	controller.CourseService.Delete(request.Context(), int64(id))

	apiResponse := helper.ApiResponseFormatter(200, "success", "create success", nil)

	helper.HandleApiResponse(writer, apiResponse)
}

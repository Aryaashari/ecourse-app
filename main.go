package main

import (
	"ecourse-app/app"
	"ecourse-app/controller"
	"ecourse-app/helper"
	"ecourse-app/repository"
	"ecourse-app/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	validate := validator.New()
	db := app.GetNewConnection()

	// Course Categories
	courseCategoryRepo := repository.NewCourseCategoryRepository()
	courseCategoryService := service.NewCourseCategoryService(courseCategoryRepo, db, validate)
	courseCategoryController := controller.NewCourseCategoryController(courseCategoryService)

	router.GET("/course/categories", courseCategoryController.FindAll)
	router.POST("/course/categories", courseCategoryController.Create)
	router.PUT("/course/categories/:courseCategoryId", courseCategoryController.Update)
	router.DELETE("/course/categories/:courseCategoryId", courseCategoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicError(&err)
}

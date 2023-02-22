package main

import (
	"ecourse-app/app"
	"ecourse-app/controller"
	"ecourse-app/exception"
	"ecourse-app/helper"
	"ecourse-app/middleware"
	"ecourse-app/repository"
	"ecourse-app/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	validator := validator.New()
	db := app.GetNewConnection()

	// Course Categories
	courseCategoryRepo := repository.NewCourseCategoryRepository()
	courseCategoryService := service.NewCourseCategoryService(courseCategoryRepo, db, validator)
	courseCategoryController := controller.NewCourseCategoryController(courseCategoryService)

	router.GET("/course/categories", courseCategoryController.FindAll)
	router.GET("/course/categories/:courseCategoryId", courseCategoryController.FindById)
	router.POST("/course/categories", courseCategoryController.Create)
	router.PUT("/course/categories/:courseCategoryId", courseCategoryController.Update)
	router.DELETE("/course/categories/:courseCategoryId", courseCategoryController.Delete)

	// Admin Authentication
	adminRepo := repository.NewAdminRepository()
	authService := service.NewAuthService(adminRepo, db, validator)
	authController := controller.NewAuthController(authService)

	router.POST("/admin/register", authController.Register)
	router.POST("/admin/login", authController.Login)

	// Courses
	courseRepository := repository.NewCourseRepository()
	courseService := service.NewCourseService(courseRepository, courseCategoryRepo, db, validator)
	courseController := controller.NewCourseController(courseService)

	router.GET("/courses", courseController.FindAll)
	router.POST("/courses", courseController.Create)
	router.GET("/courses/:courseId", courseController.FindById)
	router.PUT("/courses/:courseId", courseController.Update)
	router.DELETE("/courses/:courseId", courseController.Delete)

	// Users
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validator)
	userController := controller.NewUserController(userService)

	router.GET("/users", userController.FindAll)
	router.POST("/users", userController.Create)
	router.GET("/users/:userId", userController.FindById)
	router.PUT("/users/:userId", userController.Update)
	router.DELETE("/users/:userId", userController.Delete)

	// If has error in router
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}

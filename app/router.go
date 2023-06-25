package app

import (
	"github.com/julienschmidt/httprouter"
	"nandes007/blog-post-rest-api/controller"
	"nandes007/blog-post-rest-api/exception"
	"nandes007/blog-post-rest-api/middleware"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/users", middleware.JwtAuthMiddleware(userController.Create))
	router.GET("/api/users", middleware.JwtAuthMiddleware(userController.FindAll))
	router.POST("/api/users/login", userController.Login)
	//router.GET("/api/users", userController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}

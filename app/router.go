package app

import (
	"github.com/julienschmidt/httprouter"
	"nandes007/blog-post-rest-api/controller"
	"nandes007/blog-post-rest-api/exception"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/users", userController.Create)
	router.GET("/api/users", userController.FindAll)
	router.POST("/api/users/login", userController.Login)

	router.PanicHandler = exception.ErrorHandler

	return router
}

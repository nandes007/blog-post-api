package app

import (
	"github.com/julienschmidt/httprouter"
	"nandes007/blog-post-rest-api/controller"
	"nandes007/blog-post-rest-api/exception"
	"nandes007/blog-post-rest-api/middleware"
)

func NewRouter(userController controller.UserController, postController controller.PostController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/users", middleware.JwtAuthMiddleware(userController.Create))
	router.GET("/api/users", middleware.JwtAuthMiddleware(userController.FindAll))
	router.POST("/api/users/login", userController.Login)
	router.GET("/api/users/find", middleware.JwtAuthMiddleware(userController.Find))
	//router.GET("/api/users", userController.FindAll)

	// posts route
	router.POST("/api/posts", middleware.JwtAuthMiddleware(postController.Create))

	router.PanicHandler = exception.ErrorHandler

	return router
}

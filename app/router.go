package app

import (
	"fmt"
	"nandes007/blog-post-rest-api/controller"
	"nandes007/blog-post-rest-api/exception"
	"nandes007/blog-post-rest-api/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController, postController controller.PostController, authController controller.AuthController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Running")
	})

	// authentications
	router.POST("/api/login", authController.Login)
	router.POST("/api/register", authController.Register)

	// users route
	router.GET("/api/users", middleware.JwtAuthMiddleware(userController.FindAll))
	router.GET("/api/users/find", middleware.JwtAuthMiddleware(userController.Find))

	// posts route
	router.POST("/api/posts", middleware.JwtAuthMiddleware(postController.Create))
	router.GET("/api/posts", middleware.JwtAuthMiddleware(postController.FindAll))
	router.GET("/api/posts/:id", middleware.JwtAuthMiddleware(postController.Find))
	router.PUT("/api/posts/:id", middleware.JwtAuthMiddleware(postController.Update))
	router.DELETE("/api/posts/:id", middleware.JwtAuthMiddleware(postController.Delete))

	router.PanicHandler = exception.ErrorHandler

	return router
}

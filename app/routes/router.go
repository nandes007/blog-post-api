package routes

import (
	"fmt"
	"nandes007/blog-post-rest-api/controller"
	"nandes007/blog-post-rest-api/exception"
	"nandes007/blog-post-rest-api/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController,
	postController controller.PostController,
	authController controller.AuthController,
	commentController controller.CommentController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Running")
	})

	// authentications
	router.POST("/api/login", authController.Login)
	router.POST("/api/register", authController.Register)

	// users route
	router.GET("/api/users", middleware.JwtAuthMiddleware(userController.GetAllUsers))
	router.GET("/api/users/find", middleware.JwtAuthMiddleware(userController.GetUserByID))

	// posts route
	router.POST("/api/posts", middleware.JwtAuthMiddleware(postController.CreatePost))
	router.GET("/api/posts", middleware.JwtAuthMiddleware(postController.GetAllPosts))
	router.GET("/api/posts/:id", middleware.JwtAuthMiddleware(postController.GetPostByID))
	router.PUT("/api/posts/:id", middleware.JwtAuthMiddleware(postController.UpdatePost))
	router.DELETE("/api/posts/:id", middleware.JwtAuthMiddleware(postController.DeletePost))

	// comments route
	router.POST("/api/posts/:postId/comments", middleware.JwtAuthMiddleware(commentController.Create))

	router.PanicHandler = exception.ErrorHandler

	return router
}

package controller

import (
	"github.com/julienschmidt/httprouter"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/web"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/service"
	"net/http"
)

type PostControllerImpl struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &PostControllerImpl{
		PostService: postService,
	}
}

func (controller PostControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	postCreateRequest := post.CreateRequest{}
	helper.ReadFromRequestBody(request, &postCreateRequest)
	token := request.Header.Get("Authorization")

	postResponse := controller.PostService.Create(request.Context(), postCreateRequest, token)
	apiResponse := web.ApiResponse{
		Code:   201,
		Status: "OK",
		Data:   postResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}
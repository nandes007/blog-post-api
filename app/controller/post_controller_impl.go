package controller

import (
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/web"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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

func (controller PostControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token := request.Header.Get("Authorization")
	postsResponse := controller.PostService.FindAll(request.Context(), token)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   postsResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller PostControllerImpl) Find(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))

	helper.PanicIfError(err)

	postResponse := controller.PostService.Find(request.Context(), id)

	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller PostControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	postRequest := post.CreateRequest{}
	helper.ReadFromRequestBody(request, &postRequest)

	helper.PanicIfError(err)
	postResponse := controller.PostService.Update(request.Context(), postRequest, id)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller PostControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	controller.PostService.Delete(request.Context(), id)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

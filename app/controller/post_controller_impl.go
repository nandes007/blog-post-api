package controller

import (
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/jwt"
	"nandes007/blog-post-rest-api/model/web"
	"nandes007/blog-post-rest-api/model/web/post"
	"nandes007/blog-post-rest-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type postControllerImpl struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &postControllerImpl{
		PostService: postService,
	}
}

func (c *postControllerImpl) CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	postCreateRequest := post.PostRequest{}
	helper.ReadFromRequestBody(r, &postCreateRequest)
	token := r.Header.Get("Authorization")
	userId, err := jwt.ParseUserTokenV2(token)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.WriteToResponseBody(w, &web.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Error:  err.Error(),
		})
		return
	}

	postResponse, err := c.PostService.CreatePost(&postCreateRequest, userId)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.WriteToResponseBody(w, &web.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Error:  err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, &web.ApiResponse{
		Code:   201,
		Status: "OK",
		Data:   postResponse,
	})
}

func (c *postControllerImpl) GetAllPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	postsResponse, err := c.PostService.GetAllPosts()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.WriteToResponseBody(w, &web.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Error:  err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, &web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   postsResponse,
	})
}

func (c *postControllerImpl) GetPostByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.WriteToResponseBody(w, &web.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Error:  err.Error(),
		})
		return
	}

	postResponse, err := c.PostService.GetPostByID(id)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.WriteToResponseBody(w, &web.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Error:  err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, &web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponse,
	})
}

func (c *postControllerImpl) UpdatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.WriteToResponseBody(w, &web.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Error:  err.Error(),
		})
		return
	}

	postRequest := post.UpdatePostRequest{}
	helper.ReadFromRequestBody(r, &postRequest)
	postRequest.ID = id
	postResponse, err := c.PostService.UpdatePost(&postRequest)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.WriteToResponseBody(w, &web.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Error:  err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, &web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   postResponse,
	})
}

func (c *postControllerImpl) DeletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.WriteToResponseBody(w, &web.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Error:  err.Error(),
		})
		return
	}

	err = c.PostService.DeletePost(id)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.WriteToResponseBody(w, &web.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Error:  err.Error(),
		})
		return
	}

	helper.WriteToResponseBody(w, &web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	})
}

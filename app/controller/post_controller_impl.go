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

func (c PostControllerImpl) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	postCreateRequest := post.PostRequest{}
	helper.ReadFromRequestBody(r, &postCreateRequest)
	token := r.Header.Get("Authorization")
	postResponse, err := c.PostService.Create(r.Context(), &postCreateRequest, token)
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

func (c PostControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	token := r.Header.Get("Authorization")
	postsResponse, err := c.PostService.FindAll(r.Context(), token)
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

func (c PostControllerImpl) Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token := r.Header.Get("Authorization")
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

	postResponse, err := c.PostService.Find(r.Context(), token, id)
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

func (c PostControllerImpl) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token := r.Header.Get("Authorization")
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
	postResponse, err := c.PostService.Update(r.Context(), &postRequest, token)
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

func (c PostControllerImpl) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	err = c.PostService.Delete(r.Context(), id)
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

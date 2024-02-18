package controller

import (
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/helper/jwt"
	"nandes007/blog-post-rest-api/model/web"
	"nandes007/blog-post-rest-api/model/web/comment"
	"nandes007/blog-post-rest-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type commentControllerImpl struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &commentControllerImpl{
		commentService: commentService,
	}
}

func (c *commentControllerImpl) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	req := comment.CommentRequest{}
	helper.ReadFromRequestBody(r, &req)
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

	postId, err := strconv.Atoi(ps.ByName("postId"))
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

	response, err := c.commentService.CreateComment(&req, postId, userId)
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
		Data:   response,
	})
}

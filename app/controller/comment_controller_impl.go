package controller

import (
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/web"
	"nandes007/blog-post-rest-api/model/web/comment"
	"nandes007/blog-post-rest-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CommentControllerImpl struct {
	CommentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &CommentControllerImpl{
		CommentService: commentService,
	}
}

func (c CommentControllerImpl) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	request := comment.Request{}
	helper.ReadFromRequestBody(r, &request)
	token := r.Header.Get("Authorization")

	postId, err := strconv.Atoi(p.ByName("postId"))
	helper.PanicIfError(err)

	response := c.CommentService.Save(r.Context(), request, postId, token)

	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, apiResponse)
}

package controller

import (
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/web"
	"nandes007/blog-post-rest-api/model/web/auth"
	"nandes007/blog-post-rest-api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (c *AuthControllerImpl) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loginRequest := auth.LoginRequest{}
	helper.ReadFromRequestBody(r, &loginRequest)
	token, err := c.AuthService.Login(r.Context(), &loginRequest)
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
		Status: "Success",
		Data:   token,
	})
}

func (c *AuthControllerImpl) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	registerRequest := auth.RegisterRequest{}
	helper.ReadFromRequestBody(r, &registerRequest)
	response, err := c.AuthService.Register(r.Context(), &registerRequest)
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

package controller

import (
	"github.com/julienschmidt/httprouter"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/web"
	"nandes007/blog-post-rest-api/model/web/auth"
	"nandes007/blog-post-rest-api/service"
	"net/http"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := auth.LoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	token, err := controller.AuthService.Login(request.Context(), loginRequest)

	if err != nil {
		apiResponse := web.ApiResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   err.Error(),
		}

		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		helper.WriteToResponseBody(writer, apiResponse)
		return
	}

	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "Success",
		Data:   token,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerRequest := auth.RegisterRequest{}
	helper.ReadFromRequestBody(request, &registerRequest)
	apiResponse := web.ApiResponse{}

	response, err := controller.AuthService.Register(request.Context(), registerRequest)

	if err != nil {
		apiResponse.Code = 422
		apiResponse.Status = "Bad Request"
		apiResponse.Data = err.Error()

		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		helper.WriteToResponseBody(writer, apiResponse)
		return
	}

	apiResponse.Code = 200
	apiResponse.Status = "OK"
	apiResponse.Data = response

	helper.WriteToResponseBody(writer, apiResponse)
}

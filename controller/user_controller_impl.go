package controller

import (
	"github.com/julienschmidt/httprouter"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/web"
	"nandes007/blog-post-rest-api/model/web/user"
	"nandes007/blog-post-rest-api/service"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := user.CreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "Success",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponse := controller.UserService.FindAll(request.Context())
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "Success",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

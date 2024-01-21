package controller

import (
	"github.com/julienschmidt/httprouter"
	"nandes007/blog-post-rest-api/helper"
	"nandes007/blog-post-rest-api/model/web"
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

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponse := controller.UserService.FindAll(request.Context())
	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "Success",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

func (controller *UserControllerImpl) Find(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token := request.Header.Get("Authorization")

	user, err := controller.UserService.Find(request.Context(), token)

	if err != nil {
		apiResponse := web.ApiResponse{
			Code:   400,
			Status: "Bad Request",
			Data:   err.Error(),
		}

		helper.WriteToResponseBody(writer, apiResponse)
		return
	}

	apiResponse := web.ApiResponse{
		Code:   200,
		Status: "Success",
		Data:   user,
	}

	helper.WriteToResponseBody(writer, apiResponse)
}

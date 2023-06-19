package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

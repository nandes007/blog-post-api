package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	FindAll(write http.ResponseWriter, request *http.Request, params httprouter.Params)
	Find(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

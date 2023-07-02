package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PostController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

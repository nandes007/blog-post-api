package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

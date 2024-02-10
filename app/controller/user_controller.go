package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Find(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

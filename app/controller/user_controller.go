package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	GetUserByID(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

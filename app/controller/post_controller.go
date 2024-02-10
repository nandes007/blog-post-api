package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostController interface {
	Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

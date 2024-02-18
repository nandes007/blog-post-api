package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostController interface {
	CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	GetAllPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	GetPostByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

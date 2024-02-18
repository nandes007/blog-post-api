package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CommentController interface {
	Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

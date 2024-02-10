package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CommentController interface {
	Save(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

package middleware

import (
	"net/http"
)

type Handler struct {
	/**
	Struct for handler
	*/
	handler http.Handler
}

func NewHandler(h http.Handler) *Handler {
	return &Handler{handler: h}
}

func (handler *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.handler.ServeHTTP(w, r)
}

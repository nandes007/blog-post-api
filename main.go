package main

import (
	_ "github.com/lib/pq"
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:        "localhost:9001",
		Handler:     &myHandler{},
		ReadTimeout: 5 * time.Second,
	}

	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	mux["/tmp"] = Tmp
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "URL: "+r.URL.String())
}

func Tmp(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "version 3")
}

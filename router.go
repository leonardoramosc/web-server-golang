package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {

	handler, exist := r.FindHandler(request.URL.Path)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "<h1 style='color: red;'>Not Found</h1>")
		return
	}

	handler(w, request)

}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}

}

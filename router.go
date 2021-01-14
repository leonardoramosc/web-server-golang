package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, pathExist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, pathExist, methodExist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {

	handler, pathExist, methodExist := r.FindHandler(request.URL.Path, request.Method)

	if !pathExist {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "<h1 style='color: red;'>Not Found</h1>")
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)

}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}

}

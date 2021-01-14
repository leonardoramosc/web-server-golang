package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}

func (s *Server) Listen() error {
	fmt.Println("Server listen on port:", s.port)
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)

	if err != nil {
		return err
	}

	return nil
}

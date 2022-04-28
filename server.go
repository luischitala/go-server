package main

import (
	"net/http"
)

//Deckare the server
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

//Receiver function to allow add handlers
func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	//call the router add the method
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler

}

//Take the list of middleware to run them before execute the handler if all the middleware worked
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		//Here we pass the middleware
		f = m(f)
	}
	return f
}

//Constructor
func (s *Server) Listen() error {
	// Entry point for the app
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	} else {
		return nil
	}
}

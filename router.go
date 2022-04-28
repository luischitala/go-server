package main

import (
	"net/http"
)

// type Router struct {
// 	rules map[string]http.HandlerFunc
// }

//To add the feature of verbs, a map is pointing to another map and then points to urls
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

//Function needed to make the relation between the handler and the router
//the key is the path
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	//Very the url
	_, exist := r.rules[path]
	//then bring the handler and see if it exists for that method
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}

//Key part to run the server
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// Will use a writer w (response writer)
	// fmt.Fprintf(w, "Hello World!")
	//Find handler got the route
	handler, methodExist, exist := r.FindHandler(request.URL.Path, request.Method) //Server will know with url is calling to
	//to use both variables
	// First we will handle 404
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		//Here the return cuts the function
		return
	}
	//Verify if the method for that route is valid
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}

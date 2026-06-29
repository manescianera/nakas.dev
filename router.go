package main

import (
	"net/http"
	"strings"
)

type Router struct {
	routes  map[string]http.HandlerFunc
	static  map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]http.HandlerFunc),
		static: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) Add(path string, handler http.HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) AddStatic(path string, handler http.Handler) {
	r.static[path] = handler.ServeHTTP
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := r.routes[req.URL.Path]; ok {
		handler(w, req)
		return
	}
	for prefix, handler := range r.static {
		if strings.HasPrefix(req.URL.Path, prefix) {
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
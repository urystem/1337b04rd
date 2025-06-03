package router

import (
	"net/http"
)

func (r *router) NewServe() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", r.Catalog)
	return mux
}

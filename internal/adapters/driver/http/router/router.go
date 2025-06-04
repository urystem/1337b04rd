package router

import (
	"net/http"
)

func (r *router) NewServe() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /", r.MiddleWareInter.CheckOrSetSession(http.HandlerFunc(r.Catalog)))
	// mux.Handle("GET /", r.MiddleWareInter.CheckOrSetSession(http.HandlerFunc(r.Catalog)))

	// return r.MiddleWareInter.CheckOrSetSession(mux)
	return mux
}

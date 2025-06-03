package inbound

import "net/http"

type HandlerInter interface {
	Catalog(w http.ResponseWriter, r *http.Request)
	// CreatePost(w http.ResponseWriter, r *http.Request)
	// CreateComment(w http.ResponseWriter, r *http.Request)
	// Archive(w http.ResponseWriter, r *http.Request)
}

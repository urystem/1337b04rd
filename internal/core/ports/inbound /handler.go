package inbound

import "net/http"

type Handler interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
	CreateComment(w http.ResponseWriter, r *http.Request)
	Catalog(w http.ResponseWriter, r *http.Request)
	Archive(w http.ResponseWriter, r *http.Request)
}


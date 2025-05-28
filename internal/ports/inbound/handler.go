package inbound

import "net/http"

type HandlerInter interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
	CreateComment(w http.ResponseWriter, r *http.Request)
	Catalog(w http.ResponseWriter, r *http.Request)
	Archive(w http.ResponseWriter, r *http.Request)
}

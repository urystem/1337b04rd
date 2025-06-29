package router

import (
	"net/http"

	"1337b04rd/internal/ports/inbound"
)

// type router struct {
// 	middleware inbound.MiddleWareInter
// 	handler    inbound.HandlerInter
// }

func NewRoute(middle inbound.MiddleWareInter, hand inbound.HandlerInter) http.Handler {
	mux := http.NewServeMux()
	// mux.HandleFunc("GET /", hand.Catalog)
	mux.Handle("GET /", middle.CheckOrSetSession(http.HandlerFunc(hand.Catalog)))
	mux.HandleFunc("GET /postimage/{image}", hand.ServePostImage)
	mux.HandleFunc("GET /create-post-page", hand.CreatePostPage)
	mux.HandleFunc("GET /submit-post", hand.SubmitPost)
	// mux.Handle("GET /", r.MiddleWareInter.CheckOrSetSession(http.HandlerFunc(r.Catalog)))

	// return r.MiddleWareInter.CheckOrSetSession(mux)
	return mux
}

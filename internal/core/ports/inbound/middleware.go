package inbound

import "net/http"

type MiddleWareInter interface {
	CheckOrSetSession(next http.Handler) http.Handler
}

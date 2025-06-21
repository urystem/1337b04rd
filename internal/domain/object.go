package domain

import "io"

type Object struct {
	io.Reader
	ObjName string
	ConType string
	Size    int64
}

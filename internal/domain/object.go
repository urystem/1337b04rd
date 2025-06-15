package domain

import "io"

type Object struct {
	io.ReadCloser
	ObjName string
	ConType string
	Size    int64
}

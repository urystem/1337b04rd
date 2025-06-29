package domain

import "mime/multipart"

type PostNonContent struct {
	ID       int
	Title    string
	HasImage bool
}

type InputPost struct {
	Name    string
	Subject string
	Content string
	File    multipart.File
}

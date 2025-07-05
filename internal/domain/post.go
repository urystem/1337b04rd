package domain

import "github.com/google/uuid"

type PostNonContent struct {
	ID       uint64
	Title    string
	HasImage bool
}

type BasicInputPost struct {
	Uuid    uuid.UUID
	Name    string
	Subject string
	Content string
}

type InsertPost struct {
	BasicInputPost
	HasImage bool
}

type Form struct {
	BasicInputPost
	File *InPutObject
}

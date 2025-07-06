package domain

import (
	"time"

	"github.com/google/uuid"
)

// for catalog
type PostNonContent struct {
	ID       uint64
	Title    string
	HasImage bool
}

// create-post(like template) basic
type BasicInputPost struct {
	Uuid    uuid.UUID
	Name    string
	Subject string
	Content string
}

// create-post db
type InsertPost struct {
	BasicInputPost
	HasImage bool
}

// create-post
type Form struct {
	BasicInputPost
	File *InPutObject
}

// output
type Post struct {
	ID       uint64
	Title    string
	DataTime time.Time
}

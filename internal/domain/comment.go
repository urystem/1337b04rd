package domain

import (
	"time"

	"github.com/google/uuid"
)

// output
type Comment struct {
	ID        uint64
	UserName  string
	AvatarURL string
	ReplyToID *uint64 // nil если нет ответа на другой комментарий
	Content   string
	HasImage  bool
	DataTime  time.Time
}

// input
type BasicInputComment struct {
	PostID    uint64
	User      uuid.UUID
	Content   string
	ReplyToID *uint64 // nil если нет ответа на другой комментарий
}

// input
type CommentForm struct {
	BasicInputComment              // for sql
	File              *InPutObject // for s3
}

// sql
type InsertComment struct {
	BasicInputComment
	HasImage bool
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

type basicOutputComment struct {
	CommentID uint64
	UserName  string
	AvatarURL string
	Content   string
	HasImage  bool
	DataTime  time.Time
}

// output for archive-post like insta
type Comment struct {
	basicOutputComment
	ReplyToID *uint64 // nil если нет ответа на другой комментарий
}

// for post like tree
type CommentTree struct {
	basicOutputComment
	Replies []CommentTree // вложенные ответы
}

// input
type basicInputComment struct {
	PostID    uint64
	User      uuid.UUID
	Content   string
	ReplyToID *uint64 // nil если нет ответа на другой комментарий
}

// input
type CommentForm struct {
	basicInputComment              // for sql
	File              *InPutObject // for s3
}

// sql
type InsertComment struct {
	basicInputComment
	HasImage bool
}

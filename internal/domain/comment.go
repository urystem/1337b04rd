package domain

import "time"

type Comment struct {
	ID        uint64
	UserName  string
	AvatarURL string
	ReplyToID *uint64 // nil если нет ответа на другой комментарий
	Content   string
	HasImage  bool
	DataTime  time.Time
}

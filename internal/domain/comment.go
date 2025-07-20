package domain

type Comment struct {
	ID        uint64
	AvatarURL string
	ReplyToID *uint64 // nil если нет ответа на другой комментарий
	Content   string
	HasImage  bool
}

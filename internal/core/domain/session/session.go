package session

import (
	"time"

	"github.com/google/uuid"
)

type session struct {
	uuid      uuid.UUID
	avatarURL string
	userName  string
	// createdAt time.Time
	expiresAt time.Time
}

func NewUser() any {
	return &session{}
}

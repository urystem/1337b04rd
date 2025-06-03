package domain

import "github.com/google/uuid"

type Session struct {
	Uuid      uuid.UUID
	Name      string
	AvatarURL string
}

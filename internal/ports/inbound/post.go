package inbound

import "1337b04rd/internal/domain"

type PostInter interface {
	GetAllPosts() ([]domain.Post, error)
	SavePost(*domain.Post) error
}
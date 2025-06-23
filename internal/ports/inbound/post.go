package inbound

import "1337b04rd/internal/domain"

type PostInter interface {
	GetAllPosts() ([]domain.PostNonContent, error)
	SavePost(*domain.PostNonContent) error
}
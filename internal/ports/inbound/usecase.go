package inbound

import "1337b04rd/internal/domain"

type Usecase interface {
	ListOfPosts() ([]domain.Post, error)
}

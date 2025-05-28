package outbound

import "context"

type CharacterOutputInter interface {
	GetName() string
	GetAvatar() string
}

type CharacterInputInter interface {
	GetID() uint64
	CharacterOutputInter
}

type RickAndMortyRedisInter interface {
	SetCharacter(ctx context.Context, character CharacterInputInter) error
	GetAndDelRandomCharacter(ctx context.Context) (CharacterOutputInter, error)
}

type UseCaseRickAndMorty interface {
	GetCharacter(ctx context.Context) (CharacterOutputInter, error)
}

package outbound

import "context"

type SessionRedisInter interface {
	CheckSession(ctx context.Context, session string) (bool, error)
	SetSession(ctx context.Context, session string) error
}

type RickAndMortyRedisInter interface {
	SetCharacter(ctx context.Context, character CharacterInputInter) error
	GetAndDelRandomCharacter(ctx context.Context) (CharacterOutputInter, error)
}

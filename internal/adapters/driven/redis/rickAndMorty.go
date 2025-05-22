package redis

import (
	"context"
	"strconv"

	"1337b04rd/internal/core/ports/outbound"

	"github.com/redis/go-redis/v9"
)

type rickAndMorty struct {
	*redis.Client
}

type rickWriterReader struct {
	name      string
	avatarURL string
}

func (rick *rickAndMorty) SetCharacter(ctx context.Context, character outbound.CharacterInter) error {
	id := strconv.FormatUint(character.GetID(), 10)
	val := &rickWriterReader{character.GetName(), character.GetAvatar()}
	
	return rick.Set(ctx, id, val, 0).Err()
}

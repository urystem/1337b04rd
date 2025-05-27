package avatar

import (
	"context"

	rickandmorty "1337b04rd/internal/adapters/driven/rickAndMorty"
	"1337b04rd/internal/core/ports/outbound"
)

type rick struct {
	outbound.RickAndMortyRedisInter
}

func (rick *rick) SetCharacters(ctx context.Context) error {
	for page := 1; ; page++ {
		characters, err := rickandmorty.GetCharacters(page)
		if err != nil {
			return err
		} else if len(characters) == 0 {
			return nil
		}
		for i := range characters {
			rick.SetCharacter(ctx, characters[i])
		}
	}
}

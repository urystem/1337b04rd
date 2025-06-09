package rickandmorty

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"1337b04rd/internal/domain"
	"1337b04rd/internal/ports/outbound"
)

type rickApiClient struct {
	*http.Client
}

func InitRickApi(timeOut time.Duration) outbound.RickAndMortyApi {
	return &rickApiClient{
		&http.Client{
			Timeout: timeOut,
		},
	}
}

func (rickApi *rickApiClient) GetCharacters(ctx context.Context, p int) ([]domain.Character, error) {
	var ans domain.RickMortyReader
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character?page=%d", p)
	// req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	// if err != nil {
	// 	return nil, err
	// }
	// r, err := rickApi.client.Do(req)

	r, err := rickApi.Get(url)
	// r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&ans)
	if err != nil {
		return nil, err
	}
	return ans.Characters, nil
}

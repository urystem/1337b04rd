package rickandmorty

import (
	"encoding/json"
	"fmt"
	"net/http"

	"1337b04rd/internal/domain"
)

type rickMortyReader struct {
	Characters []domain.Character `json:"results"`
}

func GetCharacters(p int) ([]domain.Character, error) {
	var ans rickMortyReader
	url := fmt.Sprintf("https://rickandmortyapi.com/api/character?page=%d", p)
	r, err := http.Get(url)
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

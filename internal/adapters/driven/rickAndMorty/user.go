package rickandmorty

import (
	"encoding/json"
	"fmt"
	"net/http"

	"1337b04rd/internal/core/ports/outbound"
)

type rickMortyReader struct {
	Characters []character `json:"results"`
}

type character struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (c *character) GetID() uint64 {
	return c.ID
}

func (c *character) GetName() string {
	return c.Name
}

func (c *character) GetAvatar() string {
	return c.Image
}

func (r *rickMortyReader) AsInterfaceSlice() []outbound.CharacterInputInter {
	out := make([]outbound.CharacterInputInter, len(r.Characters))
	for i := range r.Characters {
		out[i] = &r.Characters[i]
	}
	return out
}

func GetCharacters(p int) ([]outbound.CharacterInputInter, error) {
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
	return ans.AsInterfaceSlice(), nil
}

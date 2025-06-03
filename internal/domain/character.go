package domain

type RickMortyReader struct {
	Characters []Character `json:"results"`
}

type Character struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

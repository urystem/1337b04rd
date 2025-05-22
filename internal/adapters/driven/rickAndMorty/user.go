package rickandmorty

import "net/http"

type rickMortyReader struct {
	characterCount uint
	pageCount      uint
	characters     []character
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

type character struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (c *character) GetID() uint {
	return c.ID
}

func (c *character) GetName() string {
	return c.Name
}

func (c *character) GetAvatar() string {
	return c.Image
}

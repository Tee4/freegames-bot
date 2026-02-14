package gamerpower

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	baseURL string
}

func NewClient() *Client {
	return &Client{
		baseURL: "https://www.gamerpower.com/api/giveaways",
	}
}

func (c *Client) Fetch() ([]Giveaway, error) {
	resp, err := http.Get(c.baseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []Giveaway
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

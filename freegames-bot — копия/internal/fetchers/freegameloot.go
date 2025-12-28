package fetchers

import (
	"encoding/json"
	"net/http"

	"freegames-bot/internal/models"
)

type FreeGameLootFetcher struct{}

func (f *FreeGameLootFetcher) Name() string {
	return "freegameloot"
}

func (f *FreeGameLootFetcher) Fetch() ([]models.Game, error) {
	resp, err := http.Get("https://freegameloot.net/api/v1/games")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []struct {
		Title string `json:"title"`
		URL   string `json:"url"`
		Store string `json:"store"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var games []models.Game

	for _, g := range data {
		games = append(games, models.Game{
			Title:   g.Title,
			EndDate: nil,
			Links: []models.GameLink{
				{
					Store: g.Store,
					URL:   g.URL,
				},
			},
		})
	}

	return games, nil
}

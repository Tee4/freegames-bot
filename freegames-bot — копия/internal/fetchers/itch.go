package fetchers

import (
	"encoding/json"
	"net/http"

	"freegames-bot/internal/models"
)

type ItchFetcher struct{}

func (f *ItchFetcher) Name() string {
	return "itch"
}

func (f *ItchFetcher) Fetch() ([]models.Game, error) {
	resp, err := http.Get("https://itch.io/games/free.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Games []struct {
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"games"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var games []models.Game

	for _, g := range data.Games {
		games = append(games, models.Game{
			Title:   g.Title,
			EndDate: nil,
			Links: []models.GameLink{
				{
					Store: "itch.io",
					URL:   g.URL,
				},
			},
		})
	}

	return games, nil
}

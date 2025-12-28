package fetchers

import (
	"encoding/json"
	"net/http"

	"freegames-bot/internal/models"
)

type GOGFetcher struct{}

func (f *GOGFetcher) Name() string {
	return "gog"
}

func (f *GOGFetcher) Fetch() ([]models.Game, error) {
	resp, err := http.Get("https://www.gog.com/games/ajax/filtered?price=free&sort=popularity")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Products []struct {
			Title string `json:"title"`
			Slug  string `json:"slug"`
		} `json:"products"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var games []models.Game

	for _, p := range data.Products {
		games = append(games, models.Game{
			Title:   p.Title,
			EndDate: nil,
			Links: []models.GameLink{
				{
					Store: "GOG",
					URL:   "https://www.gog.com/game/" + p.Slug,
				},
			},
		})
	}

	return games, nil
}

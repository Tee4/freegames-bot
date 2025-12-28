package fetchers

import (
	"encoding/json"
	"net/http"

	"freegames-bot/internal/models"
)

type SteamFetcher struct{}

func (f *SteamFetcher) Name() string {
	return "steam"
}

func (f *SteamFetcher) Fetch() ([]models.Game, error) {
	resp, err := http.Get("https://store.steampowered.com/api/featuredcategories")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Specials struct {
			Items []struct {
				Name string `json:"name"`
				URL  string `json:"store_url"`
			} `json:"items"`
		} `json:"specials"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var games []models.Game

	for _, item := range data.Specials.Items {
		games = append(games, models.Game{
			Title:   item.Name,
			EndDate: nil,
			Links: []models.GameLink{
				{
					Store: "Steam",
					URL:   item.URL,
				},
			},
		})
	}

	return games, nil
}

package fetchers

import (
	"encoding/json"
	"net/http"
	"time"

	"freegames-bot/internal/models"
)

type GamerPowerFetcher struct{}

func (f *GamerPowerFetcher) Name() string {
	return "gamerpower"
}

func (f *GamerPowerFetcher) Fetch() ([]models.Game, error) {
	resp, err := http.Get("https://www.gamerpower.com/api/giveaways")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []struct {
		Title       string `json:"title"`
		OpenGiveaway string `json:"open_giveaway_url"`
		EndDate     string `json:"end_date"`
		Platform    string `json:"platforms"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var games []models.Game

	for _, g := range data {
		var endDate *time.Time

		if g.EndDate != "" && g.EndDate != "N/A" {
			if t, err := time.Parse("2006-01-02 15:04:05", g.EndDate); err == nil {
				endDate = &t
			}
		}

		games = append(games, models.Game{
			Title:   g.Title,
			EndDate: endDate,
			Links: []models.GameLink{
				{
					Store: g.Platform,
					URL:   g.OpenGiveaway,
				},
			},
		})
	}

	return games, nil
}

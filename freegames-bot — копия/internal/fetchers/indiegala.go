package fetchers

import (
	"encoding/json"
	"net/http"

	"freegames-bot/internal/models"
)

type IndieGalaFetcher struct{}

func (f *IndieGalaFetcher) Name() string {
	return "indiegala"
}

func (f *IndieGalaFetcher) Fetch() ([]models.Game, error) {
	resp, err := http.Get("https://www.indiegala.com/get-freebies")
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
					Store: "IndieGala",
					URL:   g.URL,
				},
			},
		})
	}

	return games, nil
}

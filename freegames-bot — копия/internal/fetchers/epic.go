package fetchers

import (
	"encoding/json"
	"net/http"
	"time"

	"freegames-bot/internal/models"
)

type EpicFetcher struct{}

func (f *EpicFetcher) Name() string {
	return "epic"
}

func (f *EpicFetcher) Fetch() ([]models.Game, error) {
	resp, err := http.Get("https://store-site-backend-static.ak.epicgames.com/freeGamesPromotions")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Data struct {
			Catalog struct {
				SearchStore struct {
					Elements []struct {
						Title      string `json:"title"`
						ProductURL string `json:"productSlug"`
						Promotions struct {
							PromotionalOffers []struct {
								PromotionalOffers []struct {
									EndDate string `json:"endDate"`
								} `json:"promotionalOffers"`
							} `json:"promotionalOffers"`
						} `json:"promotions"`
					} `json:"elements"`
				} `json:"searchStore"`
			} `json:"Catalog"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var games []models.Game

	for _, el := range data.Data.Catalog.SearchStore.Elements {
		var endDate *time.Time

		if len(el.Promotions.PromotionalOffers) > 0 &&
			len(el.Promotions.PromotionalOffers[0].PromotionalOffers) > 0 {

			t, err := time.Parse(time.RFC3339, el.Promotions.PromotionalOffers[0].PromotionalOffers[0].EndDate)
			if err == nil {
				endDate = &t
			}
		}

		games = append(games, models.Game{
			Title:   el.Title,
			EndDate: endDate,
			Links: []models.GameLink{
				{
					Store: "Epic Games",
					URL:   "https://store.epicgames.com/p/" + el.ProductURL,
				},
			},
		})
	}

	return games, nil
}

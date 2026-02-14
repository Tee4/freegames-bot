package gamerpower

import (
        "fmt"
	"strings"
	"time"

	"freegames-bot/internal/domain"
)

func ToDomain(g Giveaway) domain.Giveaway {
	var endAt *time.Time
	if g.EndDate != "" && g.EndDate != "N/A" {
		if t, err := time.Parse("2006-01-02", g.EndDate); err == nil {
			endAt = &t
		}
	}

	return domain.Giveaway{
		ID:        itoa(g.ID),
		Title:     g.Title,
		URL:       g.OpenGiveawayURL,
		ImageURL:  g.Image,
		EndAt:     endAt,
		Stores:    parseStores(g.Platforms),
	}
}

func itoa(v int) string {
	return fmt.Sprintf("%d", v)
}

func parseStores(s string) []domain.Store {
	s = strings.ToLower(s)

	var stores []domain.Store
	if strings.Contains(s, "steam") {
		stores = append(stores, domain.StoreSteam)
	}
	if strings.Contains(s, "epic") {
		stores = append(stores, domain.StoreEpic)
	}
	if strings.Contains(s, "gog") {
		stores = append(stores, domain.StoreGOG)
	}

	return stores
}

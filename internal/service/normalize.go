package service

import (
	"strconv"
	"strings"
	"time"

	"freegames-bot/internal/domain"
	"freegames-bot/internal/gamerpower"
)

func Normalize(g gamerpower.Giveaway) domain.Giveaway {
	var endAt *time.Time
	if g.EndDate != "" && g.EndDate != "N/A" {
		if t, err := time.Parse("2006-01-02 15:04:05", g.EndDate); err == nil {
			endAt = &t
		}
	}

	platforms, stores := parsePlatformsAndStores(g.Platforms)

	return domain.Giveaway{
		ID:        strconv.Itoa(g.ID),
		Title:     g.Title,
		URL:       g.OpenGiveawayURL,
		ImageURL:  g.Image,
		Platforms: platforms,
		Stores:    stores,
		EndAt:     endAt,
	}
}

func NormalizeGiveaways(in []gamerpower.Giveaway) []domain.Giveaway {
	out := make([]domain.Giveaway, 0, len(in))
	for _, g := range in {
		out = append(out, Normalize(g))
	}
	return out
}

func parsePlatformsAndStores(raw string) ([]domain.Platform, []domain.Store) {
	raw = strings.ToLower(raw)

	var platforms []domain.Platform
	var stores []domain.Store

	// платформы
	if strings.Contains(raw, "pc") {
		platforms = append(platforms, domain.PlatformPC)
	}
	if strings.Contains(raw, "playstation") || strings.Contains(raw, "ps4") || strings.Contains(raw, "ps5") {
		platforms = append(platforms, domain.PlatformPS)
	}
	if strings.Contains(raw, "xbox") {
		platforms = append(platforms, domain.PlatformXbox)
	}

	// магазины
	if strings.Contains(raw, "steam") {
		stores = append(stores, domain.StoreSteam)
	}
	if strings.Contains(raw, "epic") {
		stores = append(stores, domain.StoreEpic)
	}
	if strings.Contains(raw, "gog") {
		stores = append(stores, domain.StoreGOG)
	}
	if strings.Contains(raw, "itch") {
		stores = append(stores, domain.StoreItch)
	}

	return uniquePlatforms(platforms), uniqueStores(stores)
}

func uniquePlatforms(in []domain.Platform) []domain.Platform {
	seen := map[domain.Platform]bool{}
	var out []domain.Platform
	for _, p := range in {
		if !seen[p] {
			seen[p] = true
			out = append(out, p)
		}
	}
	return out
}

func uniqueStores(in []domain.Store) []domain.Store {
	seen := map[domain.Store]bool{}
	var out []domain.Store
	for _, s := range in {
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}

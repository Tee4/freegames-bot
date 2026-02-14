package service

import "freegames-bot/internal/domain"

func FilterPlatforms(all []domain.Giveaway) []domain.Giveaway {
	out := make([]domain.Giveaway, 0, len(all))

	for _, g := range all {
		if hasAllowedPlatform(g.Platforms) {
			out = append(out, g)
		}
	}

	return out
}

func hasAllowedPlatform(p []domain.Platform) bool {
	for _, v := range p {
		switch v {
		case domain.PlatformPC,
			domain.PlatformPS,
			domain.PlatformXbox:
			return true
		}
	}
	return false
}

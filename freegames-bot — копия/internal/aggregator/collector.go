package aggregator

import (
	"strings"

	"freegames-bot/internal/models"
)

// Collect удаляет дубликаты игр по названию.
// НИ В КОЕМ СЛУЧАЕ не лезет во внутреннюю структуру Game.
func Collect(games []models.Game) []models.Game {
	seen := make(map[string]struct{})
	result := make([]models.Game, 0, len(games))

	for _, g := range games {
		title := strings.TrimSpace(g.Title)
		if title == "" {
			continue
		}

		key := strings.ToLower(title)
		if _, exists := seen[key]; exists {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, g)
	}

	return result
}

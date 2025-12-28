package fetchers

import "freegames-bot/internal/models"

type Fetcher interface {
	Name() string
	Fetch() ([]models.Game, error)
}

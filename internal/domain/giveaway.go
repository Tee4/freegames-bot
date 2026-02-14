package domain

import "time"

// ГДЕ игра запускается
type Platform string

// ОТКУДА игру забирают
type Store string

const (
	PlatformPC    Platform = "pc"
	PlatformPS    Platform = "ps"
	PlatformXbox Platform = "xbox"
)

const (
	StoreSteam Store = "steam"
	StoreEpic  Store = "epic"
	StoreGOG   Store = "gog"
	StoreItch  Store = "itch"
)

type Giveaway struct {
	ID       string
	Title    string
	URL      string
	ImageURL string

	Platforms []Platform
	Stores    []Store

	EndAt *time.Time
}

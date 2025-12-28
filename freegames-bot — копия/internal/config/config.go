package config

import "time"

type Config struct {
	CachePath string

	DiscordToken   string
	DiscordChannel string

	Interval time.Duration
}

package main

import (
	"context"
	"log"
	"os"
	"time"

	"freegames-bot/internal/cache"
	"freegames-bot/internal/config"
	"freegames-bot/internal/discord"
	"freegames-bot/internal/fetchers"

	"github.com/joho/godotenv"
)

func main() {
	// 0. load .env
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system env")
	}

	token := os.Getenv("DISCORD_TOKEN")
	channelID := os.Getenv("DISCORD_CHANNEL_ID")

	if token == "" || channelID == "" {
		log.Fatal("DISCORD_TOKEN or DISCORD_CHANNEL_ID is not set")
	}

	// 1. config
	cfg := config.Config{
		CachePath:      "cache",
		DiscordToken:   token,
		DiscordChannel: channelID,
		Interval:       6 * time.Hour,
	}

	// 2. cache
	c, err := cache.Open(cfg.CachePath)
	if err != nil {
		log.Fatal(err)
	}

	// 3. fetchers
	fs := []fetchers.Fetcher{
		fetchers.NewEpic(),
		fetchers.NewSteam(),
		fetchers.NewGOG(),
		fetchers.NewItch(),
		fetchers.NewIndieGala(),
		fetchers.NewFreeGameLoot(),
		fetchers.NewGamerPower(),
	}

	// 4. discord bot
	bot := discord.NewBot(
		cfg.DiscordToken,
		cfg.DiscordChannel,
		c,
		fs,
		cfg.Interval,
	)

	// 5. run
	ctx := context.Background()
	if err := bot.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

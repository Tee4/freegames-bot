package main

import (
	"log"
	"os"
	"strconv"
        "github.com/joho/godotenv"

	"freegames-bot/internal/gamerpower"
	"freegames-bot/internal/notify"
	"freegames-bot/internal/service"
	"freegames-bot/internal/storage"
)

func main() {
        _ = godotenv.Load()

	log.Println("freegames-bot starting")

	// ---------- ENV ----------
	token := os.Getenv("DISCORD_BOT_TOKEN")
	channelID := os.Getenv("DISCORD_CHANNEL_ID")

	if token == "" || channelID == "" {
		log.Fatal("DISCORD_BOT_TOKEN or DISCORD_CHANNEL_ID is not set")
	}

	maxPerRun := getEnvInt("MAX_PER_RUN", 10)

	// rate-limit Discord (сообщений подряд)
	discordRate := getEnvInt("DISCORD_RATE_LIMIT", 5)

	// ---------- DISCORD ----------
	discord, err := notify.NewDiscord(token, channelID, discordRate)
	if err != nil {
		log.Fatal(err)
	}

	// ---------- CLIENT ----------
	client := gamerpower.NewClient()

	// ---------- STORAGE ----------
	sentStore, err := storage.NewFileStore("data/sent.json")
	if err != nil {
		log.Fatal(err)
	}

	// переотправка через 5 дней
	deduper := service.NewDeduplicator(sentStore, 5)

	// ---------- FETCH ----------
	log.Println("fetching giveaways from GamerPower")

	raw, err := client.Fetch()
	if err != nil {
		log.Fatal(err)
	}

	// ---------- NORMALIZE ----------
	all := service.NormalizeGiveaways(raw)

	// ---------- FILTER ----------
	all = service.FilterPlatforms(all)

	// ---------- SORT ----------
	service.SortGiveaways(all)

	// ---------- SEND ----------
	sent := 0

	for _, g := range all {
		if sent >= maxPerRun {
			break
		}

		if !deduper.Allow(g) {
			continue
		}

		if err := discord.SendGiveaway(g); err != nil {
			log.Println("discord error:", err)
			continue
		}

		_ = sentStore.Mark(g.ID)
		sent++
	}

	log.Printf("done, sent=%d\n", sent)
}

// ---------- helpers ----------

func getEnvInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	if n, err := strconv.Atoi(v); err == nil {
		return n
	}
	return def
}

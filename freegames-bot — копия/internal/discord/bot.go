package discord

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"freegames-bot/internal/aggregator"
	"freegames-bot/internal/cache"
	"freegames-bot/internal/fetchers"
	"freegames-bot/internal/models"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	token     string
	channelID string
	cache     *cache.Cache
	fetchers  []fetchers.Fetcher
	interval  time.Duration
}

func NewBot(
	token string,
	channelID string,
	c *cache.Cache,
	fs []fetchers.Fetcher,
	interval time.Duration,
) *Bot {
	return &Bot{
		token:     token,
		channelID: channelID,
		cache:     c,
		fetchers:  fs,
		interval:  interval,
	}
}

func (b *Bot) Run(ctx context.Context) error {
	s, err := discordgo.New("Bot " + b.token)
	if err != nil {
		return err
	}

	if err := s.Open(); err != nil {
		return err
	}
	defer s.Close()

	log.Println("Discord bot started")

	// первый запуск сразу
	b.runOnce(s)

	ticker := time.NewTicker(b.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			b.runOnce(s)
		case <-ctx.Done():
			return nil
		}
	}
}

func (b *Bot) runOnce(s *discordgo.Session) {
	var allGames []models.Game

	// 1. FETCH
	for _, f := range b.fetchers {
		games, err := f.Fetch()
		if err != nil {
			log.Printf("[%s] fetch error: %v\n", f.Name(), err)
			continue
		}
		allGames = append(allGames, games...)
	}

	fetched := len(allGames)
	if fetched == 0 {
		log.Println("fetched=0, nothing to do")
		return
	}

	// 2. AGGREGATE (дедуп между сторами)
	aggregatedGames := aggregator.Aggregate(allGames)
	aggregated := len(aggregatedGames)

	// 3. SEND ONLY NEW
	sent := 0

	for _, g := range aggregatedGames {
		// если игра уже была — пропускаем
		if _, exists := b.cache.Get(g.Title); exists {
			continue
		}

		msg := formatGame(g)
		if msg == "" {
			// все ссылки пустые — не шлём мусор
			continue
		}

		_, err := s.ChannelMessageSend(b.channelID, msg)
		if err != nil {
			log.Println("discord send error:", err)
			continue
		}

		data, err := json.Marshal(g)
		if err != nil {
			log.Println("json marshal error:", err)
			continue
		}

		if err := b.cache.Set(g.Title, data); err != nil {
			log.Println("cache set error:", err)
			continue
		}

		sent++
	}

	log.Printf(
		"cycle done: fetched=%d aggregated=%d sent=%d\n",
		fetched,
		aggregated,
		sent,
	)
}

func formatGame(g models.Game) string {
	var sb strings.Builder
	linkCount := 0

	sb.WriteString("🎮 " + g.Title + "\n")

	for _, l := range g.Links {
		// 🔴 КРИТИЧЕСКИ ВАЖНО:
		// отбрасываем пустые / мусорные ссылки
		if strings.TrimSpace(l.URL) == "" {
			continue
		}

		sb.WriteString(fmt.Sprintf("  - %s: %s\n", l.Store, l.URL))
		linkCount++
	}

	// если нет ни одной нормальной ссылки — не шлём
	if linkCount == 0 {
		return ""
	}

	return sb.String()
}

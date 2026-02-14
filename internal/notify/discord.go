package notify

import (
	"strings"
	"time"

	"freegames-bot/internal/domain"

	"github.com/bwmarrin/discordgo"
)

type DiscordNotifier struct {
	s    *discordgo.Session
	ch   string
	rate time.Duration
}

func NewDiscord(token, channel string, rateMs int) (*DiscordNotifier, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	if err := s.Open(); err != nil {
		return nil, err
	}

	return &DiscordNotifier{
		s:    s,
		ch:   channel,
		rate: time.Millisecond * time.Duration(rateMs),
	}, nil
}

func (d *DiscordNotifier) SendGiveaway(g domain.Giveaway) error {
	embed := &discordgo.MessageEmbed{
		Title:       d.title(g),
		URL:         g.URL,
		Description: d.description(g),
		Color:       d.color(g),
		Image: &discordgo.MessageEmbedImage{
			URL: g.ImageURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Source: GamerPower",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err := d.s.ChannelMessageSendEmbed(d.ch, embed)
	time.Sleep(d.rate)
	return err
}

/* ---------- helpers ---------- */

func (d *DiscordNotifier) title(g domain.Giveaway) string {
	if isEndingSoon(g) {
		return "🔥 " + g.Title
	}
	return g.Title
}

func (d *DiscordNotifier) color(g domain.Giveaway) int {
	if isEndingSoon(g) {
		return 0xff5555 // красный
	}
	return 0x00ff99 // зелёный
}

func (d *DiscordNotifier) description(g domain.Giveaway) string {
	var b strings.Builder

	// 🎮 Платформы
	b.WriteString("🎮 **Платформы:** ")
	b.WriteString(formatPlatforms(g.Platforms))
	b.WriteString("\n")

	// ⏳ Дата окончания
	if g.EndAt != nil {
		b.WriteString("⏱️ **Доступно до:** ")
		b.WriteString(g.EndAt.Format("02 Jan 2006"))
		b.WriteString("\n")
	}

	// 🛒 Магазины
	if len(g.Stores) > 0 {
		b.WriteString("🛒 **Магазины:** ")
		b.WriteString(formatStores(g.Stores))
		b.WriteString("\n")
	}

	b.WriteString("\n👉 [Забрать игру](")
	b.WriteString(g.URL)
	b.WriteString(")")

	return b.String()
}

/* ---------- formatting ---------- */

func formatPlatforms(p []domain.Platform) string {
	if len(p) == 0 {
		return "не указано"
	}

	var names []string
	for _, platform := range p {
		switch platform {
		case domain.PlatformPC:
			names = append(names, "PC")
		case domain.PlatformPS:
			names = append(names, "PlayStation")
		case domain.PlatformXbox:
			names = append(names, "Xbox")
		}
	}

	return strings.Join(names, ", ")
}

func formatStores(stores []domain.Store) string {
	var names []string
	for _, s := range stores {
		switch s {
		case domain.StoreSteam:
			names = append(names, "Steam")
		case domain.StoreEpic:
			names = append(names, "Epic Games")
		case domain.StoreGOG:
			names = append(names, "GOG")
		case domain.StoreItch:
			names = append(names, "Itch.io")
		}
	}
	return strings.Join(names, ", ")
}

/* ---------- logic ---------- */

func isEndingSoon(g domain.Giveaway) bool {
	if g.EndAt == nil {
		return false
	}
	return time.Until(*g.EndAt) <= 72*time.Hour
}

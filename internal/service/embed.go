package service

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"freegames-bot/internal/domain"
)

func ToDiscordEmbed(g domain.Giveaway) *discordgo.MessageEmbed {
	desc := buildDescription(g)

	return &discordgo.MessageEmbed{
		Title:       g.Title,
		URL:         g.URL,
		Description: desc,
		Color:       0x00ff99,
		Image: &discordgo.MessageEmbedImage{
			URL: g.ImageURL,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Source: GamerPower",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func buildDescription(g domain.Giveaway) string {
	var b strings.Builder

	if g.EndAt != nil {
		b.WriteString("⏳ **До:** ")
		b.WriteString(g.EndAt.Format("02 Jan 2006"))
		b.WriteString("\n")
	}

	if len(g.Stores) > 0 {
		b.WriteString("🏬 **Магазины:** ")
		for i, s := range g.Stores {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(string(s))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n👉 [Забрать игру](")
	b.WriteString(g.URL)
	b.WriteString(")")

	return b.String()
}

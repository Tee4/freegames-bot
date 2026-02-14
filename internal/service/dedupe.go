package service

import (
	"log"
	"time"

	"freegames-bot/internal/domain"
	"freegames-bot/internal/storage"
)

type Deduplicator struct {
	sent        *storage.FileStore
	resendAfter time.Duration
}

func NewDeduplicator(sent *storage.FileStore, days int) *Deduplicator {
	return &Deduplicator{
		sent:        sent,
		resendAfter: time.Hour * 24 * time.Duration(days),
	}
}

func (d *Deduplicator) Allow(g domain.Giveaway) bool {
	if t, ok := d.sent.Has(g.ID); ok {
		if time.Since(t) < d.resendAfter {
			log.Printf("[dedupe] skipped sent: %s", g.ID)
			return false
		}
		log.Printf("[dedupe] resend allowed: %s", g.ID)
	}
	return true
}

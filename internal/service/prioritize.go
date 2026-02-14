package service

import (
	"sort"

	"freegames-bot/internal/domain"
)

func SortGiveaways(all []domain.Giveaway) {
	sort.SliceStable(all, func(i, j int) bool {
		a := all[i]
		b := all[j]

		// 1️⃣ Ending soon (раньше заканчивается → выше)
		if a.EndAt != nil && b.EndAt != nil {
			if !a.EndAt.Equal(*b.EndAt) {
				return a.EndAt.Before(*b.EndAt)
			}
		}

		if a.EndAt != nil && b.EndAt == nil {
			return true
		}
		if a.EndAt == nil && b.EndAt != nil {
			return false
		}

		// 2️⃣ Store priority (Steam / Epic выше)
		return storePriority(a.Stores) > storePriority(b.Stores)
	})
}

func storePriority(stores []domain.Store) int {
	for _, s := range stores {
		if s == domain.StoreSteam || s == domain.StoreEpic {
			return 2
		}
	}
	if len(stores) > 0 {
		return 1
	}
	return 0
}

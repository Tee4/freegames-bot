package storage

import "sync"

type MemoryStore struct {
	mu   sync.Mutex
	seen map[string]struct{}
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		seen: make(map[string]struct{}),
	}
}

// IsNew возвращает true, если раздача ещё не была обработана
func (s *MemoryStore) IsNew(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.seen[key]; ok {
		return false
	}

	s.seen[key] = struct{}{}
	return true
}

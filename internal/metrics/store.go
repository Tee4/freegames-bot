package metrics

import (
	"encoding/json"
	"os"
	"sync"
)

type Metrics struct {
	Sent             int `json:"sent"`
	SkippedDuplicate int `json:"skipped_duplicate"`
	SkippedQueued    int `json:"skipped_queued"`
	RateLimited      int `json:"rate_limited"`
	SendErrors       int `json:"send_errors"`
	Retries          int `json:"retries"`
}

type Store struct {
	path string
	mu   sync.Mutex
	data Metrics
}

func New(path string) (*Store, error) {
	s := &Store{path: path}
	_ = s.load()
	return s, nil
}

func (s *Store) load() error {
	b, err := os.ReadFile(s.path)
	if err != nil {
		return nil
	}
	return json.Unmarshal(b, &s.data)
}

func (s *Store) save() error {
	b, _ := json.MarshalIndent(s.data, "", "  ")
	return os.WriteFile(s.path, b, 0644)
}

func (s *Store) Inc(fn func(*Metrics)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fn(&s.data)
	_ = s.save()
}

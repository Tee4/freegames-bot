package storage

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type FileStore struct {
	mu   sync.Mutex
	path string
	data map[string]time.Time
}

func NewFileStore(path string) (*FileStore, error) {
	s := &FileStore{
		path: path,
		data: map[string]time.Time{},
	}

	b, err := os.ReadFile(path)
	if err == nil {
		_ = json.Unmarshal(b, &s.data)
	}

	return s, nil
}

func (s *FileStore) Has(id string) (time.Time, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.data[id]
	return t, ok
}

func (s *FileStore) Mark(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[id] = time.Now()

	b, _ := json.MarshalIndent(s.data, "", "  ")
	return os.WriteFile(s.path, b, 0644)
}

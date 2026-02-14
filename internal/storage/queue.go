package storage

import (
	"encoding/json"
	"os"
	"sync"
)

type Queue struct {
	mu    sync.Mutex
	path  string
	items []string
}

func NewQueue(path string) (*Queue, error) {
	q := &Queue{path: path}

	if data, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(data, &q.items)
	}

	return q, nil
}

func (q *Queue) save() error {
	data, err := json.MarshalIndent(q.items, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(q.path, data, 0644)
}

func (q *Queue) Add(id string) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.items = append(q.items, id)
	return q.save()
}

func (q *Queue) PopN(n int) []string {
	q.mu.Lock()
	defer q.mu.Unlock()

	if n > len(q.items) {
		n = len(q.items)
	}

	out := q.items[:n]
	q.items = q.items[n:]
	_ = q.save()

	return out
}

// ✅ НОВОЕ
func (q *Queue) Has(id string) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	for _, v := range q.items {
		if v == id {
			return true
		}
	}
	return false
}

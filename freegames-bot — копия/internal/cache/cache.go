package cache

import (
	"database/sql"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Cache struct {
	db *sql.DB
}

func Open(dir string) (*Cache, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	dbPath := filepath.Join(dir, "cache.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS games (
		title TEXT PRIMARY KEY,
		data  BLOB,
		updated INTEGER
	);
	`)
	if err != nil {
		return nil, err
	}

	return &Cache{db: db}, nil
}

func (c *Cache) Get(title string) ([]byte, bool) {
	row := c.db.QueryRow(`SELECT data FROM games WHERE title = ?`, title)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false
	}
	return data, true
}

func (c *Cache) Set(title string, data []byte) error {
	_, err := c.db.Exec(
		`INSERT OR REPLACE INTO games(title, data, updated) VALUES (?, ?, ?)`,
		title, data, time.Now().Unix(),
	)
	return err
}

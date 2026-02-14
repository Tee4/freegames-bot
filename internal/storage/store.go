package storage

type Store interface {
	IsNew(key string) bool
}

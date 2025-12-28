package lock

import (
	"fmt"
	"os"
)

type Lock struct {
	path string
	file *os.File
}

func Acquire(path string) (*Lock, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return nil, fmt.Errorf("lock exists")
	}

	return &Lock{
		path: path,
		file: f,
	}, nil
}

func (l *Lock) Release() {
	if l.file != nil {
		l.file.Close()
		_ = os.Remove(l.path)
	}
}

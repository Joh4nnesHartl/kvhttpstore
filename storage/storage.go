package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	storagePath string
)

func init() {
	SetStoragePath("files")

	os.MkdirAll(storagePath, 0644)
}

// KVStorage represents an in memory key-value storage
type KVStorage struct{}

// Store stores the value at key
func (s KVStorage) Store(key string, value []byte) error {
	filename := filepath.Join(storagePath, key)

	err := ioutil.WriteFile(filename, value, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Receive returns the value stored at key
func (s KVStorage) Receive(key string) (value []byte, ok bool) {
	filename := filepath.Join(storagePath, key)

	value, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, false
	}

	return value, true
}

// SetStoragePath sets the Path where the files are stored
func SetStoragePath(name string) {
	wd, _ := os.Getwd()

	storagePath = filepath.Join(wd, name)
}

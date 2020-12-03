package storage

// KVStorage represents an in memory key-value storage
type KVStorage map[string][]byte

// Store stores the value at key
func (s KVStorage) Store(key string, value []byte) {
	s[key] = value
}

// Receive returns the value stored at key
func (s KVStorage) Receive(key string) (value []byte, ok bool) {
	value, ok = s[key]

	return value, ok
}

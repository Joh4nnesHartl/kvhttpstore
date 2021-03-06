package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	wd, _ := os.Getwd()
	storagePath = filepath.Join(wd, "testdata")

	os.MkdirAll(storagePath, 0644)
}

func TestStoreRecieve(t *testing.T) {
	testCases := []struct {
		desc  string
		key   string
		value []byte
	}{
		{
			"0",
			"submarine",
			[]byte{0x69, 0x42, 0x00},
		},
		{
			"1",
			"hello123",
			[]byte{0x00, 0x01},
		},
		{
			"2",
			"test12",
			[]byte{0x05, 0x04, 0x03, 0x02, 0x01, 0x00},
		},
	}

	var storage KVStorage

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			storage.Store(tC.key, tC.value)

			value, ok := storage.Receive(tC.key)
			require.True(t, ok)

			assert.Equal(t, value, tC.value)
		})
	}

	os.RemoveAll(storagePath)
}

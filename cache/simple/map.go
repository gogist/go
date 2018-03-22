package cache

import (
	"errors"
	"time"
)

type Map struct {
	cache      map[string][]byte
	cacheTimer map[string]time.Time
	expiration time.Duration
}

// NewMap creates a new cache using golang default type map.
// Parameter expiration is used to set the expiration of each value.
// If expiration is less than or equal to 0, then no expiration will be set.
func NewMap(expiration time.Duration) *Map {
	if expiration <= 0 {
		expiration = 0
	}
	return &Map{
		cache:      map[string][]byte{},
		cacheTimer: map[string]time.Time{},
		expiration: expiration,
	}
}

// Contains returns whether the key is in the cache or not.
// Contains will return false for an expired key.
func (m *Map) Contains(key string) bool {
	if t, ok := m.cacheTimer[key]; ok {
		if m.expiration > 0 {
			if t.Add(m.expiration).After(time.Now()) {
				return true
			}
		} else {
			return true
		}
	}
	return false
}

// Expiration returns the expiration time. If no expiration is set,
// 0 will be returned.
func (m *Map) Expiration() time.Duration {
	return m.expiration
}

// Set sets a value to a key. For the default map cache, error will be nil all
// the time.
func (m *Map) Set(key string, val []byte) error {
	m.cache[key] = val
	m.cacheTimer[key] = time.Now()
	return nil
}

// Get gets the value for a key. If the key is not in the map, an error will be
// returned.
func (m *Map) Get(key string) ([]byte, error) {
	if m.Contains(key) {
		return m.cache[key], nil
	}
	return []byte{}, errors.New(key + " does not exist")
}

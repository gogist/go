package cache

import (
	"time"
)

type Cache interface {
	Expiration() time.Duration
	Contains(string) bool
	Set(string, []byte) error
	Get(string) ([]byte, error)
}

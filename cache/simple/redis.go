package cache

import (
	"time"

	"gopkg.in/redis.v3"
)

type Redis struct {
	redisClient *redis.Client
	expiration  time.Duration
}

// NewRedis creates a new cache using redis database.
// Parameter addr, passwd are the server address and password of redis server.
// Parameter db specify the db number.
// Parameter expiration is used to set the expiration of each value.
// If expiration is less than or equal to 0, then no expiration will be set.
func NewRedis(addr, passwd string, db int64, expiration time.Duration) *Redis {
	if expiration <= 0 {
		expiration = 0
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	return &Redis{
		redisClient: redisClient,
		expiration:  expiration,
	}
}

// Contains returns whether the key is in the cache or not.
// Contains will return false for an expired key.
func (r *Redis) Contains(key string) bool {
	if re, err := r.redisClient.Exists(key).Result(); err == nil {
		return re
	}
	return false
}

// Expiration returns the expiration time. If no expiration is set,
// 0 will be returned.
func (r *Redis) Expiration() time.Duration {
	return r.expiration
}

// Set sets a value to a key.
func (r *Redis) Set(key string, val []byte) error {
	return r.redisClient.Set(key, val, r.expiration).Err()
}

// Get gets the value for a key. If the key is not in the cache, an error will be
// returned.
func (r *Redis) Get(key string) ([]byte, error) {
	return r.redisClient.Get(key).Bytes()
}

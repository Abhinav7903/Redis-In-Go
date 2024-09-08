package idis

import "time"

type Repository interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	Exists(key string) bool
	Expire(key string, ttl time.Duration) error
	TTL(key string) (time.Duration, error)
}

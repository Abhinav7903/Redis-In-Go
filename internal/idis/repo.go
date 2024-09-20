package idis

import "time"

type Repository interface {
	Set(key string, values ...string) error
	Get(key string) ([]string, error)
	Delete(key string) error
	Exists(key string) bool
	Expire(key string, ttl time.Duration) error
	TTL(key string) (time.Duration, error)
	RandomValues(key string, offset int) ([]string, error)
	SetUnique(key string, values ...string) error
	RemoveValue(key string, value string) error
	GetUnique(key string) ([]string, error)
	GetKeyFromValue(value string) ([]string, error)
	DumpToFile(filename string) error
	LoadFromDump(filename string) error
}

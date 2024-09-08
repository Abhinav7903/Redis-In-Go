package idis

import (
	"errors"
	"sync"
	"time"
)

type InMemoryRepository struct {
	store  map[string]string
	mu     sync.RWMutex
	expiry map[string]time.Time
}

// NewInMemoryRepository creates a new instance of InMemoryRepository
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		store:  make(map[string]string),
		expiry: make(map[string]time.Time),
	}
}

// Set adds or updates a key-value pair in the store
func (r *InMemoryRepository) Set(key string, value string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[key] = value
	return nil
}

// Get retrieves a value based on the given key
func (r *InMemoryRepository) Get(key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	value, ok := r.store[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value, nil
}

// Delete removes a key-value pair from the store
func (r *InMemoryRepository) Delete(key string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.store[key]; !ok {
		return errors.New("key not found")
	}
	delete(r.store, key)
	return nil
}

// Exists checks if a key exists in the store
func (r *InMemoryRepository) Exists(key string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if expiration, ok := r.expiry[key]; ok && time.Now().After(expiration) {
		return false // Key has expired
	}
	_, exists := r.store[key]
	return exists
}

func (r *InMemoryRepository) Expire(key string, ttl time.Duration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.store[key]; !ok {
		return errors.New("key not found")
	}
	r.expiry[key] = time.Now().Add(ttl)
	return nil
}

// TTL returns the remaining time-to-live for a key
func (r *InMemoryRepository) TTL(key string) (time.Duration, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	expiration, ok := r.expiry[key]
	if !ok {
		return -1, errors.New("no TTL set or key not found")
	}

	if time.Now().After(expiration) {
		delete(r.store, key)
		delete(r.expiry, key)
		return -1, errors.New("key has expired")
	}

	return time.Until(expiration), nil
}

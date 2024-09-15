package idis

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

type InMemoryRepository struct {
	store         map[string][]string
	mu            sync.RWMutex
	expiry        map[string]time.Time
	reverseLookup map[string][]string // Map value to a slice of keys
}

// NewInMemoryRepository creates a new instance of InMemoryRepository
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		store:         make(map[string][]string),
		expiry:        make(map[string]time.Time),
		reverseLookup: make(map[string][]string),
	}
}

// Set adds one or more values to a key (appends values to the key's slice)
func (r *InMemoryRepository) Set(key string, values ...string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// If the key already exists, append the new values
	if existingValues, ok := r.store[key]; ok {
		r.store[key] = append(existingValues, values...)
	} else {
		// If the key doesn't exist, create a new slice with the values
		r.store[key] = values
	}

	// Update reverse lookup map
	for _, value := range values {
		r.reverseLookup[value] = append(r.reverseLookup[value], key)
	}

	return nil
}

// Get retrieves all values associated with a key
func (r *InMemoryRepository) Get(key string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	values, ok := r.store[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return values, nil
}

// Delete removes a key and its associated values from the store
func (r *InMemoryRepository) Delete(key string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if existingValues, ok := r.store[key]; ok {
		for _, value := range existingValues {
			// Remove the key from the list of keys for each value
			keys := r.reverseLookup[value]
			for i, k := range keys {
				if k == key {
					r.reverseLookup[value] = append(keys[:i], keys[i+1:]...)
					break
				}
			}
			if len(r.reverseLookup[value]) == 0 {
				delete(r.reverseLookup, value)
			}
		}
		delete(r.store, key)
		delete(r.expiry, key)
		return nil
	}

	return errors.New("key not found")
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

// Expire sets the expiration time for a key
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

// RandomValues returns a specific number of random values from the key's associated list
func (r *InMemoryRepository) RandomValues(key string, count int) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	values, ok := r.store[key]
	if !ok {
		return nil, errors.New("key not found")
	}

	// Check if the requested count is valid
	if count <= 0 || count > len(values) {
		return nil, errors.New("invalid count value")
	}

	// Seed the random number generator for randomness
	rand.Seed(time.Now().UnixNano())

	// Shuffle the values slice
	shuffledValues := make([]string, len(values))
	copy(shuffledValues, values)
	rand.Shuffle(len(shuffledValues), func(i, j int) {
		shuffledValues[i], shuffledValues[j] = shuffledValues[j], shuffledValues[i]
	})

	// Return the requested number of random values
	return shuffledValues[:count], nil
}

// SetUnique adds unique values to a key, ensuring no duplicates
func (r *InMemoryRepository) SetUnique(key string, values ...string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Use a map to track unique values
	uniqueValues := make(map[string]bool)

	// Load existing values into the map if the key already exists
	if existingValues, ok := r.store[key]; ok {
		for _, value := range existingValues {
			uniqueValues[value] = true
		}
	}

	// Add new values to the map
	for _, value := range values {
		uniqueValues[value] = true
	}

	// Convert map to slice
	uniqueSlice := make([]string, 0, len(uniqueValues))
	for value := range uniqueValues {
		uniqueSlice = append(uniqueSlice, value)
	}

	// Remove old values from reverse lookup map
	if existingValues, ok := r.store[key]; ok {
		for _, value := range existingValues {
			if _, exists := uniqueValues[value]; !exists {
				// Remove the key from the list of keys for the value
				keys := r.reverseLookup[value]
				for i, k := range keys {
					if k == key {
						r.reverseLookup[value] = append(keys[:i], keys[i+1:]...)
						break
					}
				}
				if len(r.reverseLookup[value]) == 0 {
					delete(r.reverseLookup, value)
				}
			}
		}
	}

	// Store updated unique values
	r.store[key] = uniqueSlice

	// Update reverse lookup map
	for _, value := range uniqueSlice {
		r.reverseLookup[value] = append(r.reverseLookup[value], key)
	}

	return nil
}

// RemoveValue removes a specific value from a key
func (r *InMemoryRepository) RemoveValue(key string, value string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	values, ok := r.store[key]
	if !ok {
		return errors.New("key not found")
	}

	for i, v := range values {
		if v == value {
			// Remove the value from reverse lookup map
			keys := r.reverseLookup[value]
			for j, k := range keys {
				if k == key {
					r.reverseLookup[value] = append(keys[:j], keys[j+1:]...)
					break
				}
			}
			if len(r.reverseLookup[value]) == 0 {
				delete(r.reverseLookup, value)
			}

			// Remove from key's values
			r.store[key] = append(values[:i], values[i+1:]...)
			return nil
		}
	}

	return errors.New("value not found")
}

// GetUnique retrieves all unique values associated with a key
func (r *InMemoryRepository) GetUnique(key string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	values, ok := r.store[key]
	if !ok {
		return nil, errors.New("key not found")
	}

	// Use a map to track unique values
	valueSet := make(map[string]bool)
	var uniqueValues []string

	for _, value := range values {
		if !valueSet[value] {
			uniqueValues = append(uniqueValues, value)
			valueSet[value] = true
		}
	}

	return uniqueValues, nil
}

// GetKeyFromValue retrieves all keys associated with a specific value
func (r *InMemoryRepository) GetKeyFromValue(value string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	keys, ok := r.reverseLookup[value]
	if !ok {
		return nil, errors.New("value not found")
	}

	return keys, nil
}

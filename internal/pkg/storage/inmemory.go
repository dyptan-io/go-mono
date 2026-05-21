// Package storage implements data storage types for the persistence layer.
package storage

import (
	"errors"
	"sync"
)

var (
	ErrMissingID = errors.New("missing record ID")
	ErrNotFound  = errors.New("record not found")
)

type (
	// ID is the unique identifier for a stored record.
	ID string
	// Record is the constraint for types stored in InMemory.
	Record interface {
		ID() ID
	}
)

// Matcher filters records in a Find call.
type Matcher[T Record] func(T) bool

// InMemory is a generic, concurrency-safe in-memory store.
type InMemory[T Record] struct {
	mu      sync.RWMutex
	records map[ID]T
}

// NewInMemory returns an empty InMemory store for type T.
func NewInMemory[T Record]() *InMemory[T] {
	return &InMemory[T]{mu: sync.RWMutex{}, records: make(map[ID]T)}
}

// Get returns the record with the given ID or ErrNotFound.
func (s *InMemory[T]) Get(id ID) (T, error) {
	if id == "" {
		var zero T
		return zero, ErrMissingID
	}

	s.mu.RLock()
	r, ok := s.records[id]
	s.mu.RUnlock()

	if !ok {
		return r, ErrNotFound
	}

	return r, nil
}

// Find returns all records for which match returns true.
func (s *InMemory[T]) Find(match Matcher[T]) ([]T, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var res []T

	for _, v := range s.records {
		if match(v) {
			res = append(res, v)
		}
	}

	return res, nil
}

// Insert stores r, overwriting any existing record with the same ID.
func (s *InMemory[T]) Insert(r T) error {
	if r.ID() == "" {
		return ErrMissingID
	}

	s.mu.Lock()
	s.records[r.ID()] = r
	s.mu.Unlock()

	return nil
}

// Delete removes the record with the given ID or returns ErrNotFound.
func (s *InMemory[T]) Delete(id ID) error {
	if id == "" {
		return ErrMissingID
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.records[id]; !ok {
		return ErrNotFound
	}

	delete(s.records, id)

	return nil
}

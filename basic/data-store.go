package main

import (
	"errors"
	"sync"
)

var (
	ErrEmptyKey     = errors.New("key cannot be empty")
	ErrAlreadyExist = errors.New("key already exist")
	ErrKeyNotFound  = errors.New("key does not exist")
)

// Generic DataStore with any string keys, and generic values
type DataStore[V any] struct {
	mu    sync.RWMutex
	store map[string]V
}

func NewDataStore[V any]() *DataStore[V] {
	return &DataStore[V]{
		store: make(map[string]V),
	}
}

func (ds *DataStore[V]) Insert(key string, value V) error {

	if key == "" {
		return ErrEmptyKey
	}

	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, found := ds.store[key]; !found {
		ds.store[key] = value
		return nil
	}

	return ErrAlreadyExist
}

func (ds *DataStore[V]) Read(key string) (V, error) {
	var zero V

	if key == "" {
		return zero, ErrEmptyKey
	}

	ds.mu.RLock()
	defer ds.mu.RUnlock()
	if val, found := ds.store[key]; found {
		return val, nil
	}

	return zero, ErrKeyNotFound
}

func (ds *DataStore[V]) Remove(key string) error {

	if key == "" {
		return ErrEmptyKey
	}

	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, found := ds.store[key]; found {
		delete(ds.store, key)
		return nil
	}

	return ErrKeyNotFound
}

func (ds *DataStore[V]) Count() int {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return len(ds.store)
}

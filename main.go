package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrEmptyKey     = errors.New("key cannot be empty")
	ErrAlreadyExist = errors.New("key already exist")
	ErrKeyNotFound  = errors.New("key does not exist")
)

func main() {

	// create a new data store
	ds := NewDataStore[string]()
	total := 100
	for i := range total {
		err := ds.Insert(fmt.Sprintf("key-%d", i), fmt.Sprintf("val-%d", i))
		if err != nil {
			fmt.Printf("Expected no error, got error: %v\n", err)
		}
	}

	// read key's from data store
	for i := range total {
		key := fmt.Sprintf("key-%d", i)
		val, err := ds.Read(key)
		if err != nil {
			fmt.Printf("Expected no error, got error: %v\n", err)
		}
		fmt.Printf("key: %s, val: %s\n", key, val)
	}

	// total count from data store
	count := ds.Count()
	fmt.Println("total items in the data store: ", count)
}

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

package main

import (
	"context"
	"fmt"
	"sync"
)

type mapRepository[V any] struct {
	mu    sync.RWMutex
	store map[string]V
}

func NewMapRepository[V any]() *mapRepository[V] {
	return &mapRepository[V]{
		store: make(map[string]V),
	}
}

func (mr *mapRepository[V]) Insert(key string, val V) error {

	if key == "" {
		return ErrEmptyKey
	}

	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, found := mr.store[key]; !found {
		mr.store[key] = val
		return nil
	}

	return ErrAlreadyExist
}

func (mr *mapRepository[V]) Read(ctx context.Context, key string) (V, error) {

	var zero V
	if key == "" {
		return zero, ErrEmptyKey
	}

	mr.mu.RLock()
	defer mr.mu.RUnlock()

	select {
	case <-ctx.Done():
		if err := ctx.Err(); err != nil {
			return zero, fmt.Errorf("context error, %v", err)
		}
	default:
		if val, found := mr.store[key]; found {
			return val, nil
		}
	}

	return zero, ErrKeyNotFound
}

func (mr *mapRepository[V]) Remove(key string) error {

	if key == "" {
		return ErrEmptyKey
	}

	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, found := mr.store[key]; found {
		delete(mr.store, key)
		return nil
	}

	return ErrKeyNotFound
}

func (mr *mapRepository[V]) Count() int {
	mr.mu.RLock()
	defer mr.mu.Unlock()

	return len(mr.store)
}

func (mr *mapRepository[V]) Close() error {
	return nil
}

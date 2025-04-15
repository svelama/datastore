package main

import (
	"context"
	"errors"
)

type DataStore[V any] struct {
	repo Repository[V]
}

var (
	ErrEmptyKey     = errors.New("key cannot be empty")
	ErrAlreadyExist = errors.New("key already exist")
	ErrKeyNotFound  = errors.New("key does not exist")
)

func NewDataStore[V any](repo Repository[V]) *DataStore[V] {
	return &DataStore[V]{
		repo: repo,
	}
}

// Repository defines the storage operations independent of the underlying implementations
type Repository[V any] interface {
	Insert(key string, val V) error
	Read(ctx context.Context, key string) (V, error)
	Remove(key string) error
	Count() int
	Close() error
}

func (ds *DataStore[V]) Insert(key string, val V) error {
	return ds.repo.Insert(key, val)
}

func (ds *DataStore[V]) Read(ctx context.Context, key string) (V, error) {
	return ds.repo.Read(ctx, key)
}

func (ds *DataStore[V]) Remove(key string) error {
	return ds.repo.Remove(key)
}

func (ds *DataStore[V]) Count() int {
	return ds.repo.Count()
}

func (ds *DataStore[V]) Close() error {
	return ds.repo.Close()
}

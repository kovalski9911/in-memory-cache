package inmemorycache

import (
	"context"
	"time"
)

type item[T any] struct {
	Value  T
	cancel context.CancelFunc
}

type cache[T any] struct {
	storage map[string]item[T]
}

func New[T any]() *cache[T] {
	return &cache[T]{
		storage: make(map[string]item[T]),
	}
}

func (c *cache[T]) Set(k string, v T, ttl int) {
	duration := time.Duration(ttl) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), duration)

	go func() {
		select {
		case <-ctx.Done():
			c.Delete(k)
		}
	}()

	c.storage[k] = item[T]{
		Value:  v,
		cancel: cancel,
	}
}

func (c *cache[T]) Get(k string) T {
	return c.storage[k].Value
}

func (c *cache[T]) Delete(k string) {
	item, exists := c.storage[k]
	if exists {
		item.cancel()
		delete(c.storage, k)
	}

}

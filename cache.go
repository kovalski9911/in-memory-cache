package inmemorycache

type cache[T any] struct {
	storage map[string]T
}

func New[T any]() *cache[T] {
	return &cache[T]{
		storage: make(map[string]T),
	}
}

func (c *cache[T]) Set(k string, v T) {
	c.storage[k] = v
}

func (c *cache[T]) Get(k string) T {
	return c.storage[k]
}

func (c *cache[T]) Delete(k string) {
	delete(c.storage, k)
}

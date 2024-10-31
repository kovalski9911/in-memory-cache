package inmemorycache

type cache struct {
	storage map[string]int
}

func New() *cache {
	return &cache{
		storage: make(map[string]int),
	}
}

func (c *cache) Set(k string, v int) {
	c.storage[k] = v
}

func (c *cache) Get(k string) int {
	return c.storage[k]
}

func (c *cache) Delete(k string) {
	delete(c.storage, k)
}

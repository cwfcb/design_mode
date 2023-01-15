package strategy

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	size         int
	capacity     int
}

func NewCache(algo EvictionAlgo, capacity int) *Cache {
	return &Cache{
		storage:      make(map[string]string),
		evictionAlgo: algo,
		capacity:     capacity,
	}
}
func (c *Cache) SetEvictionAlgo(algo EvictionAlgo) {
	c.evictionAlgo = algo
}

func (c *Cache) Get(key string) string {
	value := c.storage[key]
	c.evictionAlgo.adjustKey(c.storage, key)
	return value
}

func (c *Cache) Add(key, value string) {
	if c.size == c.capacity {
		c.evict()
	}
	c.size++
	c.storage[key] = value
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c.storage)
	c.size--
}

package strategy

import "testing"

func TestCache(t *testing.T) {
	lru := &LRUAlgo{}

	cache := NewCache(lru, 2)
	cache.Add("1", "1")

	cache.SetEvictionAlgo(&LFUAlgo{})
	cache.Add("2", "2")

	cache.SetEvictionAlgo(&FIFOAlgo{})
	cache.Add("3", "3")

	val := cache.Get("3")
	t.Log(val)

}

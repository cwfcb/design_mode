package strategy

import "fmt"

type EvictionAlgo interface {
	evict(storage map[string]string)
	adjustKey(storage map[string]string, key string)
}

type LRUAlgo struct {
}

func (L *LRUAlgo) adjustKey(storage map[string]string, key string) {
	fmt.Printf("adjust key: %s by lru strategy\n", key)
}

func (L *LRUAlgo) evict(storage map[string]string) {
	fmt.Println("Evicting by lru strategy")
}

type LFUAlgo struct {
}

func (L *LFUAlgo) adjustKey(storage map[string]string, key string) {
	fmt.Printf("adjust key: %s by lfu strategy\n", key)
}

func (L *LFUAlgo) evict(storage map[string]string) {
	fmt.Println("Evicting by lfu strategy")
}

type FIFOAlgo struct {
}

func (F *FIFOAlgo) adjustKey(storage map[string]string, key string) {
	fmt.Printf("adjust key: %s by fifo strategy\n", key)
}

func (F *FIFOAlgo) evict(storage map[string]string) {
	fmt.Println("Evicting by fifo strategy")
}

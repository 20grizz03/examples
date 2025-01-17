package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type CacheAIPredict interface {
	Get() int
}

type Cache struct {
	cachedValue int
	mu          sync.RWMutex
}

func NewCache() *Cache {
	c := &Cache{
		cachedValue: aiPredict(),
	}
	go c.update()
	return c

}

func (c *Cache) Get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cachedValue
}

func (c *Cache) update() {
	for _ = range time.Tick(1 * time.Second) {
		c.mu.Lock()
		c.cachedValue = aiPredict()
		c.mu.Unlock()
	}
}

func aiPredict() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(100)
}

func main() {
	c := NewCache1()
	http.HandleFunc("/predict", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"result\": %d}", c.Get())
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

// Atomic======================================================>
type Cache1 struct {
	cachedValue atomic.Int64
}

func NewCache1() *Cache1 {
	c := &Cache1{}
	c.cachedValue.Store(int64(aiPredict()))
	go c.update1()
	return c

}

func (c *Cache1) Get() int {
	return int(c.cachedValue.Load())
}

func (c *Cache1) update1() {
	for _ = range time.Tick(1 * time.Second) {
		c.cachedValue.Store(int64(aiPredict()))
	}
}

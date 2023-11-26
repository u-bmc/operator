// SPDX-License-Identifier: BSD-3-Clause

package ipc

import (
	"context"
	"sync"
	"time"
)

type CacheItem struct {
	Value  []byte
	Expiry time.Time
	Read   bool
}

type Cache struct {
	m               sync.Map
	read            sync.Map
	lifetime        time.Duration
	maxValuesPerKey int
	topics          map[string]*sync.Map
}

func NewCache(ctx context.Context, lifetime time.Duration, maxValuesPerKey int) *Cache {
	c := &Cache{
		m:               sync.Map{},
		read:            sync.Map{},
		lifetime:        lifetime,
		maxValuesPerKey: maxValuesPerKey,
		topics:          make(map[string]*sync.Map),
	}

	go c.cleanup(ctx)

	return c
}

func (c *Cache) Write(topic, key string, val []byte) bool {
	topicMap, exists := c.topics[topic]
	if !exists {
		topicMap = &sync.Map{}
		c.topics[topic] = topicMap
	}

	item := &CacheItem{
		Value:  val,
		Expiry: time.Now().Add(c.lifetime),
		Read:   false,
	}

	oldItemsInterface, loaded := topicMap.LoadOrStore(key, []*CacheItem{item})
	if loaded {
		oldItems := oldItemsInterface.([]*CacheItem)
		if len(oldItems) >= c.maxValuesPerKey {
			return false
		}
		topicMap.Store(key, append(oldItems, item))
	}

	return true
}

func (c *Cache) Read(topic, key string) ([]byte, bool) {
	topicMap, exists := c.topics[topic]
	if !exists {
		return []byte{}, false
	}

	val, ok := topicMap.Load(key)
	if !ok {
		return []byte{}, false
	}

	items := val.([]*CacheItem)
	if len(items) == 0 {
		return []byte{}, false
	}

	item := items[0]
	item.Read = true
	return item.Value, true
}

func (c *Cache) cleanup(ctx context.Context) {
	ticker := time.NewTicker(c.lifetime)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			for _, topicMap := range c.topics {
				topicMap.Range(func(key, value interface{}) bool {
					items := value.([]*CacheItem)
					if len(items) > 0 && items[0].Read && time.Now().After(items[0].Expiry) {
						topicMap.Store(key, items[1:])
					}
					return true
				})
			}
		}
	}
}

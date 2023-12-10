// SPDX-License-Identifier: BSD-3-Clause

package cache

import (
	"context"
	"sync"
	"time"
)

type Item struct {
	Value  []byte
	Expiry time.Time
	Read   bool
}

type Cache struct {
	m               sync.Map
	read            sync.Map
	lifetime        time.Duration
	maxValuesPerKey int
}

func NewCache(ctx context.Context, lifetime time.Duration, maxValuesPerKey int) *Cache {
	c := &Cache{
		m:               sync.Map{},
		read:            sync.Map{},
		lifetime:        lifetime,
		maxValuesPerKey: maxValuesPerKey,
	}

	go c.cleanup(ctx)

	return c
}

// Write adds a new item to the cache.
// It takes a key and a value as parameters.
// The function returns a boolean indicating whether the operation was successful.
// If the number of items for the given key has reached the maximum limit, the oldest item is replaced.
func (c *Cache) Write(key string, val []byte) {
	// Create a new Item with the given value and a calculated expiry time
	item := &Item{
		Value:  val,
		Expiry: time.Now().Add(c.lifetime),
		Read:   false,
	}

	// Try to load the old items for the given key
	// If the key does not exist, store the new item in the keyMap
	oldItemsInterface, loaded := c.m.LoadOrStore(key, []*Item{item})
	if loaded {
		// If the key exists, check if the number of items has reached the maximum limit
		oldItems := oldItemsInterface.([]*Item)
		if len(oldItems) >= c.maxValuesPerKey {
			// If the limit is reached, remove the oldest item and append the new item
			oldItems = append(oldItems[1:], item)
		} else {
			// If the limit is not reached, append the new item to the old items
			oldItems = append(oldItems, item)
		}
		// Store the updated items in the keyMap
		c.m.Store(key, oldItems)
	}
}

// Read retrieves the first unread item from the cache for a given topic and key.
// It returns the item's value and a boolean indicating whether the operation was successful.
func (c *Cache) Read(key string) ([]byte, bool) {
	// Load the items for the given key
	itemsInterface, exists := c.m.Load(key)
	if !exists {
		// If the key does not exist, return an empty byte slice and false
		return []byte{}, false
	}

	// Convert the loaded items to a slice of *Item
	items := itemsInterface.([]*Item)
	if len(items) == 0 {
		// If there are no items, return an empty byte slice and false
		return []byte{}, false
	}

	// Get the first item and mark it as read
	item := items[0]
	item.Read = true
	// Return the item's value and true
	return item.Value, true
}

// cleanup periodically removes read items from the cache that have expired.
func (c *Cache) cleanup(ctx context.Context) {
	// Create a new ticker that triggers every lifetime duration
	ticker := time.NewTicker(c.lifetime)
	defer ticker.Stop()

	// Start an infinite loop
	for {
		select {
		case <-ctx.Done():
			// If the context is done, return from the function
			return
		case <-ticker.C:
			// If the ticker has triggered, iterate over all keys in the cache
			c.m.Range(func(key, value interface{}) bool {
				// Convert the value to a slice of *Item
				items := value.([]*Item)
				// If the first item has been read and its expiry time has passed, remove it from the slice
				if len(items) > 0 && items[0].Read && time.Now().After(items[0].Expiry) {
					c.m.Store(key, items[1:])
				}
				// Continue the iteration
				return true
			})
		}
	}
}

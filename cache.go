package cache

import "time"

type Cache struct {
	storage map[string]cacheEntry
}

type cacheEntry struct {
	value    string
	deadline time.Time
}

func (copy cacheEntry) isExpired() bool {
	if copy.deadline.IsZero() || time.Until(copy.deadline) > 0 {
		return false
	} else {
		return true
	}
}

func NewCache() Cache {
	return Cache{storage: make(map[string]cacheEntry)}
}

func (self *Cache) Get(key string) (string, bool) {
	entry, ok := self.storage[key]
	if ok && !entry.isExpired() {
		return entry.value, ok
	} else {
		return "", false
	}
}

func (self *Cache) Put(key, value string) {
	self.storage[key] = cacheEntry{value: value}
}

func (self *Cache) Keys() []string {
	keys := make([]string, 0, len(self.storage))
	for key, entry := range self.storage {
		if !entry.isExpired() {
			keys = append(keys, key)
		}
	}
	return keys
}

func (self *Cache) PutTill(key, value string, deadline time.Time) {
	self.storage[key] = cacheEntry{value: value, deadline: deadline}
}

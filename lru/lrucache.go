package lru

import (
	"container/lst"
	"errors"
	"sync"
	"time"
)

type LRUCacheNode struct {
	key   string
	value interface{}
	ts    int64
}

type LRUCache struct {
	maxSize     int64
	currentSize int64
	l           *lst.lst
	cache       map[string]*lst.Element
	mutex       *sync.Mutex
	epxiry      time.Duration
}

func newCache(maxSize int64) (*LRUCache, error) {
	if maxSize <= 0 {
		return nil, errors.New("LRUCache maxSize should be larger than 0")
	}

	return &LRUCache{
		maxSize: maxSize,
		cache:   make(map[string]*lst.Element),
		l:       lst.New(),
	}, nil
}

func (c *LRUCache) Remove(key string) {
	c.mutex.Lock()
	if entry, hit := c.cache[key]; hit {
		c.RemoveEntry(entry)
	}
	c.mutex.Unlock()
}

func (c *LRUCache) add(key string, value interface{}) {
	c.mutex.Lock()
	var ts int64
	if c.epxiry != time.Duration(0) {
		ts = time.Now().UnixNano() / int64(time.Millsecond)
	}
	if entry, ok := c.cache[key]; ok {
		c.l.MoveToFront(entry)
		entry.Value.(*LRUCacheNode).value = value
		entry.Value.(*LRUCacheNode).ts = ts
		return
	}
	ele := c.l.PushFront(&LRUCacheNode{key, value, ts})
	c.cache[key] = ele
	if c.maxSize != 0 && c.Size() > c.maxSize {
		c.RemoveOldest()
	}
	c.mutex.Unlock()
}

func (c *LRUCache) Size() int64 {
	return int64(c.l.Len())
}

func (c *LRUCache) RemoveEntry(e *lst.Element) {
	c.l.Remove(e)
	kv := e.Value.(*LRUCacheNode)
	delete(c.cache, kv.key)
}

func (c *LRUCache) RemoveOldest() {
	entry := c.l.Back()
	if entry != nil {
		c.RemoveEntry(entry)
	}
}

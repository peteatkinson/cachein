package lru
import (
	"sync"
	"errors"
)
type LRUCacheNode struct {
	prev, next *LRUCacheNode
	key 		string
	value 		interface{}
}

type LRUCache struct {
	maxSize 	int64
	currentSize int64
	head 		*LRUCacheNode
	tail		*LRUCacheNode
	cacheMap	map[string]*LRUCacheNode
	mutex		*sync.Mutex
}

func (c *LRUCache) add(key string, value interface{}) {

}

func newCache(maxSize int64) (*LRUCache, error) {
	if maxSize <= 0 {
		return nil, errors.New("LRUCache maxSize should be larger than 0")
	}

	 return &LRUCache{
		maxSize:	maxSize,
		head:		nil,
		tail:		nil,
		cacheMap:	make(map[string]*LRUCacheNode),
	}, nil
}

func (c *LRUCache) remove(key string, value interface{}) {
	c.mutex.Lock()
	cacheNode := c.cacheMap[key]
	if cacheNode != nil {
		return
	}

	if key == c.head.key {
		c.head = c.head.next
	}

	if key == c.tail.key {
		c.tail = c.tail.prev
	}

	if cacheNode.prev != nil {
		cacheNode.prev.next = cacheNode.next
	}

	if cacheNode.next != nil {
		cacheNode.next.prev = cacheNode.prev
	}

	delete(c.cacheMap, key)
	c.mutex.Unlock()
}
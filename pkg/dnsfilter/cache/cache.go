package cache

import (
	"sync"
	"time"

	"github.com/miekg/dns"
)

type Element struct {
	Answer    *dns.RR
	timeAdded time.Time
}

type Cache struct {
	elements map[string]Element
	mutex    sync.RWMutex
}

const (
	expirationTime = 20 * time.Minute
)

func New() *Cache {
	return &Cache{
		elements: make(map[string]Element, 30),
	}
}

func (c *Cache) Get(d string) (Element, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	element, ok := c.elements[d]
	if !ok {
		return Element{}, false
	}

	if time.Since(element.timeAdded) > expirationTime {
		delete(c.elements, d)
		return Element{}, false
	}

	return element, true
}

func (c *Cache) Set(d string, answer *dns.RR) {
	c.mutex.Lock()

	c.elements[d] = Element{
		timeAdded: time.Now(),
		Answer:    answer,
	}

	c.mutex.Unlock()
}

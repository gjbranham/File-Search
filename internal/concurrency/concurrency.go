package concurrency

import (
	"sync"
)

type FileInfo struct {
	Key, File string
	LineNum   int
}

type MatchInfo struct {
	mu      sync.Mutex
	Count   int
	Matches []FileInfo
}

func (c *MatchInfo) CounterInc(count int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Count = c.Count + count
}

func (c *MatchInfo) AddMatch(info ...FileInfo) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Matches = append(c.Matches, info...)
}

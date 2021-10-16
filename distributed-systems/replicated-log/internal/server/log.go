package server

import (
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read() ([]Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.records, nil
}

// START:types
type Record struct {
	Value  string `json:"value"`
	Offset uint64 `json:"offset"`
}

//END:types

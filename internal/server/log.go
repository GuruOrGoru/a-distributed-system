package server

import (
	"fmt"
	"sync"
)

type Log struct {
	Records []Record
	mu      sync.Mutex
}

func NewLog() *Log {
	return &Log{}
}

type Record struct {
	Value  []byte
	Offset uint64
}

var ErrOffsetNotFound = fmt.Errorf("Error the given offset is not found")

func (l *Log) Append(r Record) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	r.Offset = uint64(len(l.Records))
	l.Records = append(l.Records, r)

	return r.Offset, nil
}

func (l *Log) Read(offset uint64) (Record, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if offset >= uint64(len(l.Records)) {
		return Record{}, ErrOffsetNotFound
	}

	return l.Records[offset], nil
}

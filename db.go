// Package db provides a key-value store.
package db

import "sync"

type DB struct {
	mu    sync.RWMutex
	state map[string]string
}

func New() *DB {
	return &DB{state: make(map[string]string)}
}

func (d *DB) Set(k, v string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.state[k] = v
}

func (d *DB) Get(k string) (string, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	v, ok := d.state[k]
	return v, ok
}

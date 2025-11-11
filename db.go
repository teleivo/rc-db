// Package db provides a key-value store.
package db

type DB struct {
	state map[string]string
}

func New() *DB {
	return &DB{state: make(map[string]string)}
}

func (d *DB) Set(k, v string) {
	d.state[k] = v
}

func (d *DB) Get(k string) (string, bool) {
	v, ok := d.state[k]
	return v, ok
}

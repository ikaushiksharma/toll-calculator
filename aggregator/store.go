package main

import "github.com/ikaushiksharma/toll-calculator/types"

type MemoryStore struct {}
func(m *MemoryStore) Insert(d types.Distance) error {
	return nil
}
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}	
	
}

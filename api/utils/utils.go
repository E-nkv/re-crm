package utils

import "sync"

type Object map[string]any

type void struct{}

type Set[T comparable] map[T]void

func (s *Set[T]) Contains(key T) bool {
	_, exists := (*s)[key]
	return exists
}
func (s *Set[T]) Add(key T) {
	(*s)[key] = void{}
}

type ThreadSafeSet[T comparable] struct {
	set Set[T]
	mu  sync.Mutex
}

func (tsSet *ThreadSafeSet[T]) Contains(key T) bool {
	tsSet.mu.Lock()
	defer tsSet.mu.Unlock()
	return tsSet.set.Contains(key)
}
func (tsSet *ThreadSafeSet[T]) Add(key T) {
	tsSet.mu.Lock()
	defer tsSet.mu.Unlock()
	tsSet.set.Add(key)
}

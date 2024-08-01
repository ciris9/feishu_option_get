package set

import (
	"cmp"
	"sync"
)

type Set[T cmp.Ordered] struct {
	mu sync.RWMutex
	m  map[T]struct{}
}

func (s *Set[T]) Add(data T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.m[data]; !ok {
		s.m[data] = struct{}{}
	}
}

func (s *Set[T]) Range(fun func(value T) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k := range s.m {
		if fun(k) {
			continue
		} else {
			break
		}
	}
}

func (s *Set[T]) Remove(data T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, data)
}

func (s *Set[T]) IsExist(data T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.m[data]
	return ok
}

func (s *Set[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.m)
}

func New[T cmp.Ordered]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

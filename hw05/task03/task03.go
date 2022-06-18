package task03

import "sync"

type Set struct {
	sync.Mutex
	mm map[float64]struct{}
}

func NewSet() *Set {
	return &Set{
		mm: map[float64]struct{}{},
	}
}

func (s *Set) Add(i float64) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *Set) Has(i float64) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

type SetRW struct {
	sync.RWMutex
	mm map[float64]struct{}
}

func NewSetRW() *SetRW {
	return &SetRW{
		mm: map[float64]struct{}{},
	}
}

func (s *SetRW) AddRW(i float64) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *SetRW) HasRW(i float64) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}

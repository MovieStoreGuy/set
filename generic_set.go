package set

import (
	"fmt"
	"reflect"
	"sync"
)

// New creates concrete set type
// that is thread safe
func New(vals ...interface{}) Set {
	s := &setImpl{
		set: make(map[interface{}]struct{}, 0),
	}
	for _, i := range vals {
		s.Insert(i)
	}
	return s
}

type setImpl struct {
	rw  sync.RWMutex
	set map[interface{}]struct{}
}

func (s *setImpl) Insert(item interface{}) {
	switch reflect.ValueOf(item).Kind() {
	case reflect.Array, reflect.Slice, reflect.Map:
		panic("unsupported kind used")
	}
	if a, ok := item.(Set); ok {
		for _, i := range a.Items() {
			s.Insert(i)
		}
		return
	}
	s.rw.Lock()
	defer s.rw.Unlock()
	s.set[item] = struct{}{}
}

func (s *setImpl) HasItem(item interface{}) bool {
	s.rw.RLock()
	defer s.rw.RUnlock()
	_, exist := s.set[item]
	return exist
}

func (s *setImpl) Items() []interface{} {
	s.rw.RLock()
	defer s.rw.RUnlock()
	items := make([]interface{}, 0, len(s.set))
	for item := range s.set {
		items = append(items, item)
	}
	return items
}

func (s *setImpl) Size() int {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return len(s.set)
}

func (s *setImpl) String() string {
	s.rw.RLock()
	defer s.rw.RUnlock()
	return fmt.Sprint(s.Items())
}

package knex

import "reflect"

type typeSet struct {
	set map[reflect.Type]bool
}

func newTypeSet() *typeSet {
	return &typeSet{make(map[reflect.Type]bool)}
}

func (s *typeSet) add(i reflect.Type) bool {
	_, found := s.set[i]
	s.set[i] = true
	return !found
}

func (s *typeSet) get(i reflect.Type) bool {
	_, found := s.set[i]
	return found
}

func (s *typeSet) remove(i reflect.Type) {
	delete(s.set, i)
}

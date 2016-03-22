package knex

import "reflect"

type typeSet struct {
	set map[reflect.Type]bool
}

func newTypeSet() *typeSet {
	return &typeSet{make(map[reflect.Type]bool)}
}

func (self *typeSet) add(i reflect.Type) bool {
	_, found := self.set[i]
	self.set[i] = true
	return !found
}

func (self *typeSet) get(i reflect.Type) bool {
	_, found := self.set[i]
	return found
}

func (self *typeSet) remove(i reflect.Type) {
	delete(self.set, i)
}

package utility

import "reflect"

type typeSet struct {
	set map[reflect.Type]bool
}

func newTypeSet() *typeSet {
	return &typeSet{make(map[reflect.Type]bool)}
}

func (set *typeSet) add(i reflect.Type) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found
}

func (set *typeSet) get(i reflect.Type) bool {
	_, found := set.set[i]
	return found
}

func (set *typeSet) remove(i reflect.Type) {
	delete(set.set, i)
}

package main

// Set struct will be used to declare the Set type object
type Set struct {
	// We will use map to implement the set.
	// To make map work as a set the type of value in a map is bool.
	hashMap map[int]bool
}

// NewSet will initialize and return the new object of Set.
func NewSet() *Set {
	s := new(Set)
	s.hashMap = make(map[int]bool)
	return s
}

// Add will add the value in the Set.
func (s *Set) Add(value int) {
	s.hashMap[value] = true
}

// Delete will delete the value from the set.
func (s *Set) Delete(value int) {
	delete(s.hashMap, value)
}

// Exists will check if the value exists in the set or not.
func (s *Set) Exists(value int) bool {
	_, ok := s.hashMap[value]
	return ok
}

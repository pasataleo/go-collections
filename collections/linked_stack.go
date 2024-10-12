package collections

import (
	"iter"

	"github.com/pasataleo/go-objects/objects"
)

type linkedStack[O objects.Object] struct {
	list *linkedList[O]
}

// NewStack creates a new stack.
func NewStack[O objects.Object]() Stack[O] {
	return &linkedStack[O]{
		list: &linkedList[O]{
			first: nil,
			last:  nil,
			size:  0,
		},
	}
}

// Object implementation

// Equals implements objects.Object.
func (s *linkedStack[O]) Equals(other any) bool {
	if otherQ, ok := other.(Stack[O]); ok {
		return s.list.Equals(otherQ)
	}
	return false
}

// HashCode implements objects.Object.
func (s *linkedStack[O]) HashCode() uint64 {
	return s.list.HashCode() * 13997
}

// String implements objects.Object.
func (s *linkedStack[O]) String() string {
	return s.list.String()
}

// MarshalJSON implements objects.Object.
func (s *linkedStack[O]) MarshalJSON() ([]byte, error) {
	return s.list.MarshalJSON()
}

// UnmarshalJSON implements objects.Object.
func (s *linkedStack[O]) UnmarshalJSON(bytes []byte) error {
	return s.list.UnmarshalJSON(bytes)
}

// Iterable implementation

// Iterator implements objects.Iterable.
func (s *linkedStack[O]) Iterator() objects.Iterator[O] {
	return s.list.Iterator()
}

// Collection implementation

// Elems implements objects.Collection.
func (s *linkedStack[O]) Elems() iter.Seq[O] {
	return objects.SequenceFrom[O](s)
}

// Add implements objects.Collection.
func (s *linkedStack[O]) Add(value O) error {
	return s.Offer(value)
}

// AddAll implements objects.Collection.
func (s *linkedStack[O]) AddAll(values Collection[O]) error {
	return collectionAddAll[O](s, values)
}

// Remove implements objects.Collection.
func (s *linkedStack[O]) Remove(value O) error {
	return s.list.Remove(value)
}

// RemoveAll implements objects.Collection.
func (s *linkedStack[O]) RemoveAll(values Collection[O]) error {
	return collectionRemoveAll[O](s, values)
}

// Contains implements objects.Collection.
func (s *linkedStack[O]) Contains(value O) bool {
	return s.list.Contains(value)
}

// ContainsAll implements objects.Collection.
func (s *linkedStack[O]) ContainsAll(values Collection[O]) bool {
	return collectionContainsAll[O](s, values)
}

// Copy implements objects.Collection.
func (s *linkedStack[O]) Copy() Collection[O] {
	return &linkedStack[O]{
		list: s.list.Copy().(*linkedList[O]),
	}
}

// Size implements objects.Collection.
func (s *linkedStack[O]) Size() int {
	return s.list.Size()
}

// IsEmpty implements objects.Collection.
func (s *linkedStack[O]) IsEmpty() bool {
	return s.list.Size() == 0
}

// Clear implements objects.Collection.
func (s *linkedStack[O]) Clear() {
	s.list.Clear()
}

// Stack implementation

// Offer implements Stack.
func (s *linkedStack[O]) Offer(value O) error {
	return s.list.Insert(value, s.list.size)
}

// Peep implements Stack.
func (s *linkedStack[O]) Peep() (O, error) {
	return s.list.Get(s.list.size - 1)
}

// Pop implements Stack.
func (s *linkedStack[O]) Pop() (O, error) {
	return s.list.RemoveAt(s.list.size - 1)
}

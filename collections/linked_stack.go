package collections

import "github.com/pasataleo/go-objects/objects"

type linkedStack[O any] struct {
	list *linkedList[O]
}

func NewStack[O objects.Object]() Stack[O] {
	return &linkedStack[O]{
		list: &linkedList[O]{
			first:     nil,
			last:      nil,
			size:      0,
			converter: objects.ObjectIdentityConverter[O](),
		},
	}
}

func NewStackT[O any](converter objects.ObjectConverter[O]) Stack[O] {
	return &linkedStack[O]{
		list: &linkedList[O]{
			first:     nil,
			last:      nil,
			size:      0,
			converter: converter,
		},
	}
}

// Object implementation

func (s *linkedStack[O]) Equals(other any) bool {
	if otherQ, ok := other.(Stack[O]); ok {
		return s.list.Equals(otherQ)
	}
	return false
}

func (s *linkedStack[O]) HashCode() uint64 {
	return s.list.HashCode() * 13997
}

func (s *linkedStack[O]) ToString() string {
	return s.list.ToString()
}

// Iterable implementation

func (s *linkedStack[O]) Iterator() objects.Iterator[O] {
	return s.list.Iterator()
}

// Collection implementation

func (s *linkedStack[O]) Add(value O) error {
	return s.Offer(value)
}

func (s *linkedStack[O]) AddAll(values Collection[O]) error {
	return collectionAddAll[O](s, values)
}

func (s *linkedStack[O]) Remove(value O) error {
	return s.list.Remove(value)
}

func (s *linkedStack[O]) RemoveAll(values Collection[O]) error {
	return collectionRemoveAll[O](s, values)
}

func (s *linkedStack[O]) Contains(value O) bool {
	return s.list.Contains(value)
}

func (s *linkedStack[O]) ContainsAll(values Collection[O]) bool {
	return collectionContainsAll[O](s, values)
}

func (s *linkedStack[O]) Copy() Collection[O] {
	return &linkedStack[O]{
		list: s.list.Copy().(*linkedList[O]),
	}
}

func (s *linkedStack[O]) Size() int {
	return s.list.Size()
}

func (s *linkedStack[O]) IsEmpty() bool {
	return s.list.Size() == 0
}

// Stack implementation

func (s *linkedStack[O]) Offer(value O) error {
	return s.list.Insert(value, s.list.size)
}

func (s *linkedStack[O]) Peep() (O, error) {
	return s.list.Get(s.list.size - 1)
}

func (s *linkedStack[O]) Pop() (O, error) {
	return s.list.RemoveAt(s.list.size - 1)
}

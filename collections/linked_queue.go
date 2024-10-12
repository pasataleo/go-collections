package collections

import (
	"iter"

	"github.com/pasataleo/go-objects/objects"
)

type linkedQueue[O objects.Object] struct {
	list *linkedList[O]
}

// NewQueue creates a new queue.
func NewQueue[O objects.Object]() Queue[O] {
	return &linkedQueue[O]{
		list: &linkedList[O]{
			first: nil,
			last:  nil,
			size:  0,
		},
	}
}

// Object implementation

// Equals implements objects.Object.
func (q *linkedQueue[O]) Equals(other any) bool {
	return queueEquals[O](q, other)
}

// HashCode implements objects.Object.
func (q *linkedQueue[O]) HashCode() uint64 {
	return queueHashCode[O](q)
}

// String implements objects.Object.
func (q *linkedQueue[O]) String() string {
	return q.list.String()
}

// MarshalJSON implements objects.Object.
func (q *linkedQueue[O]) MarshalJSON() ([]byte, error) {
	return q.list.MarshalJSON()
}

// UnmarshalJSON implements objects.Object.
func (q *linkedQueue[O]) UnmarshalJSON(bytes []byte) error {
	return q.list.UnmarshalJSON(bytes)
}

// Iterable implementation

// Iterator implements objects.Iterable.
func (q *linkedQueue[O]) Iterator() objects.Iterator[O] {
	return q.list.Iterator()
}

// Collection implementation

// Elems implements objects.Collection.
func (q *linkedQueue[O]) Elems() iter.Seq[O] {
	return objects.SequenceFrom[O](q)
}

// Add implements objects.Collection.
func (q *linkedQueue[O]) Add(value O) error {
	return q.Offer(value)
}

// AddAll implements objects.Collection.
func (q *linkedQueue[O]) AddAll(values Collection[O]) error {
	return collectionAddAll[O](q, values)
}

// Remove implements objects.Collection.
func (q *linkedQueue[O]) Remove(value O) error {
	return q.list.Remove(value)
}

// RemoveAll implements objects.Collection.
func (q *linkedQueue[O]) RemoveAll(values Collection[O]) error {
	return collectionRemoveAll[O](q, values)
}

// Contains implements objects.Collection.
func (q *linkedQueue[O]) Contains(value O) bool {
	return q.list.Contains(value)
}

// ContainsAll implements objects.Collection.
func (q *linkedQueue[O]) ContainsAll(values Collection[O]) bool {
	return collectionContainsAll[O](q, values)
}

// Copy implements objects.Collection.
func (q *linkedQueue[O]) Copy() Collection[O] {
	return &linkedQueue[O]{
		list: q.list.Copy().(*linkedList[O]),
	}
}

// Size implements objects.Collection.
func (q *linkedQueue[O]) Size() int {
	return q.list.Size()
}

// IsEmpty implements objects.Collection.
func (q *linkedQueue[O]) IsEmpty() bool {
	return q.list.Size() == 0
}

// Clear implements objects.Collection.
func (q *linkedQueue[O]) Clear() {
	q.list.Clear()
}

// Queue implementation

// Offer adds an element to the queue.
func (q *linkedQueue[O]) Offer(value O) error {
	return q.list.Insert(value, q.list.size)
}

// Peep returns the first element in the queue.
func (q *linkedQueue[O]) Peep() (O, error) {
	return q.list.Get(0)
}

// Pop removes and returns the first element in the queue.
func (q *linkedQueue[O]) Pop() (O, error) {
	return q.list.RemoveAt(0)
}

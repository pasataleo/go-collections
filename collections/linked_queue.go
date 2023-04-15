package collections

import "github.com/pasataleo/go-objects/objects"

type linkedQueue[O any] struct {
	list *linkedList[O]
}

func NewQueue[O objects.Object]() Queue[O] {
	return &linkedQueue[O]{
		list: &linkedList[O]{
			first:     nil,
			last:      nil,
			size:      0,
			converter: objects.ObjectIdentityConverter[O](),
		},
	}
}

func NewQueueT[O any](converter objects.ObjectConverter[O]) Queue[O] {
	return &linkedQueue[O]{
		list: &linkedList[O]{
			first:     nil,
			last:      nil,
			size:      0,
			converter: converter,
		},
	}
}

// Object implementation

func (q *linkedQueue[O]) Equals(other any) bool {
	return queueEquals(q, other, q.list.converter)
}

func (q *linkedQueue[O]) HashCode() uint64 {
	return queueHashCode[O](q, q.list.converter)
}

func (q *linkedQueue[O]) ToString() string {
	return q.list.ToString()
}

// Iterable implementation

func (q *linkedQueue[O]) Iterator() objects.Iterator[O] {
	return q.list.Iterator()
}

// Collection implementation

func (q *linkedQueue[O]) Add(value O) error {
	return q.Offer(value)
}

func (q *linkedQueue[O]) AddAll(values Collection[O]) error {
	return collectionAddAll[O](q, values)
}

func (q *linkedQueue[O]) Remove(value O) error {
	return q.list.Remove(value)
}

func (q *linkedQueue[O]) RemoveAll(values Collection[O]) error {
	return collectionRemoveAll[O](q, values)
}

func (q *linkedQueue[O]) Contains(value O) bool {
	return q.list.Contains(value)
}

func (q *linkedQueue[O]) ContainsAll(values Collection[O]) bool {
	return collectionContainsAll[O](q, values)
}

func (q *linkedQueue[O]) Copy() Collection[O] {
	return &linkedQueue[O]{
		list: q.list.Copy().(*linkedList[O]),
	}
}

func (q *linkedQueue[O]) Size() int {
	return q.list.Size()
}

func (q *linkedQueue[O]) IsEmpty() bool {
	return q.list.Size() == 0
}

// Queue implementation

func (q *linkedQueue[O]) Offer(value O) error {
	return q.list.Insert(value, q.list.size)
}

func (q *linkedQueue[O]) Peep() (O, error) {
	return q.list.Get(0)
}

func (q *linkedQueue[O]) Pop() (O, error) {
	return q.list.RemoveAt(0)
}

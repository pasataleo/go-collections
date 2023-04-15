package collections

import (
	"bytes"
	"fmt"

	"github.com/pasataleo/go-errors/errors"
	"github.com/pasataleo/go-objects/objects"
)

type heap[O any] struct {
	items []O

	converter  objects.ObjectConverter[O]
	comparator objects.Comparator[O]
}

func NewPriorityQueue[O objects.ComparableObject[O]]() Queue[O] {
	return NewPriorityQueueT(objects.ObjectIdentityConverter[O](), objects.ComparableComparator[O]())
}

func NewPriorityQueueO[O objects.Object](comparator objects.Comparator[O]) Queue[O] {
	return NewPriorityQueueT(objects.ObjectIdentityConverter[O](), comparator)
}

func NewPriorityQueueC[O objects.Comparable[O]](converter objects.ObjectConverter[O]) Queue[O] {
	return NewPriorityQueueT(converter, objects.ComparableComparator[O]())
}

func NewPriorityQueueT[O any](converter objects.ObjectConverter[O], comparator objects.Comparator[O]) Queue[O] {
	return &heap[O]{
		items:      nil,
		converter:  converter,
		comparator: comparator,
	}
}

// Object implementation

func (h *heap[O]) Equals(other any) bool {
	return queueEquals(h, other, h.converter)
}

func (h *heap[O]) HashCode() uint64 {
	return queueHashCode[O](h, h.converter)
}

func (h *heap[O]) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")

	first := true
	for iterator := h.Iterator(); iterator.HasNext(); {
		value := iterator.Next()
		if first {
			buffer.WriteString(h.converter.ToString(value))
		} else {
			buffer.WriteString(fmt.Sprintf(",%s", h.converter.ToString(value)))
		}
		first = false
	}
	buffer.WriteString("]")
	return buffer.String()
}

// Iterable implementation

type heapIterator[O any] struct {
	safe *heap[O]
}

func (h *heapIterator[O]) HasNext() bool {
	return h.safe.Size() > 0
}

func (h *heapIterator[O]) Next() O {
	value, err := h.safe.Pop()
	if err != nil {
		panic(err)
	}
	return value
}

func (h *heap[O]) Iterator() objects.Iterator[O] {
	return &heapIterator[O]{
		safe: h.Copy().(*heap[O]),
	}
}

// Collection implementation

func (h *heap[O]) Add(value O) error {
	return h.Offer(value)
}

func (h *heap[O]) AddAll(values Collection[O]) error {
	return collectionAddAll[O](h, values)
}

func (h *heap[O]) Remove(value O) error {
	for ix, current := range h.items {
		if h.converter.Equals(current, value) {
			_, err := h.remove(ix)
			return err
		}
	}
	return errors.Embed(errors.New(nil, ErrorCodeNotFound, "not found"), value)
}

func (h *heap[O]) RemoveAll(values Collection[O]) error {
	return collectionRemoveAll[O](h, values)
}

func (h *heap[O]) Contains(value O) bool {
	for _, item := range h.items {
		if h.converter.Equals(item, value) {
			return true
		}
	}
	return false
}

func (h *heap[O]) ContainsAll(values Collection[O]) bool {
	return collectionContainsAll[O](h, values)
}

func (h *heap[O]) Copy() Collection[O] {
	var contents []O
	for _, value := range h.items {
		contents = append(contents, value)
	}
	return &heap[O]{
		items:      contents,
		converter:  h.converter,
		comparator: h.comparator,
	}
}

func (h *heap[O]) Size() int {
	return len(h.items)
}

func (h *heap[O]) IsEmpty() bool {
	return len(h.items) == 0
}

// Queue implementation

func (h *heap[O]) Offer(value O) error {
	ix := len(h.items)
	h.items = append(h.items, value)
	h.up(ix)
	return nil
}

func (h *heap[O]) Peep() (O, error) {
	if len(h.items) == 0 {
		var null O
		return null, errors.New(nil, ErrorCodeOutOfBounds, "out of bounds")
	}

	return h.items[0], nil
}

func (h *heap[O]) Pop() (O, error) {
	return h.remove(0)
}

// heap implementation

func (h *heap[O]) remove(ix int) (O, error) {
	if len(h.items) == 0 {
		var null O
		return null, errors.New(nil, ErrorCodeOutOfBounds, "out of bounds")
	}

	value := h.items[ix]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.down(ix)
	return value, nil

}

func (h *heap[O]) up(ix int) {
	if ix == 0 {
		return
	}

	parent := (ix - 1) / 2
	cmp := h.comparator.Compare(h.items[ix], h.items[parent])
	if cmp < 0 {
		h.swap(ix, parent)
		h.up(parent)
	}
}

func (h *heap[O]) down(ix int) {
	if ix >= len(h.items) {
		return
	}

	var child int
	one := (ix * 2) + 1
	two := (ix * 2) + 2

	if one >= len(h.items) {
		return
	} else if two >= len(h.items) {
		child = one
	} else {
		cmp := h.comparator.Compare(h.items[one], h.items[two])
		if cmp < 0 {
			child = one
		} else {
			child = two
		}
	}

	cmp := h.comparator.Compare(h.items[ix], h.items[child])
	if cmp > 0 {
		h.swap(ix, child)
		h.down(child)
	}
}

func (h *heap[O]) swap(ix, jx int) {
	value := h.items[ix]
	h.items[ix] = h.items[jx]
	h.items[jx] = value
}

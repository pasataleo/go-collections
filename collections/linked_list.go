package collections

import (
	"github.com/pasataleo/go-errors/errors"
	"github.com/pasataleo/go-objects/objects"
)

type linkedList[O any] struct {
	first *linkedListNode[O]
	last  *linkedListNode[O]

	size int

	converter objects.ObjectConverter[O]
}

type linkedListNode[O any] struct {
	before *linkedListNode[O]
	after  *linkedListNode[O]

	value O
}

// Object implementation

func (list *linkedList[O]) Equals(other any) bool {
	return listEquals[O](list, other, list.converter)
}

func (list *linkedList[O]) HashCode() uint64 {
	return listHashCode[O](list, list.converter)
}

func (list *linkedList[O]) ToString() string {
	return listString[O](list, list.converter)
}

// Iterable implementation

func (list *linkedList[O]) Iterator() objects.Iterator[O] {
	return &linkedListIterator[O]{
		current: list.first,
	}
}

// Collection implementation

func (list *linkedList[O]) Contains(value O) bool {
	return list.IndexOf(value) >= 0
}

func (list *linkedList[O]) ContainsAll(values Collection[O]) bool {
	return collectionContainsAll[O](list, values)
}

func (list *linkedList[O]) Add(value O) error {
	return list.Insert(value, list.size)
}

func (list *linkedList[O]) AddAll(values Collection[O]) error {
	return collectionAddAll[O](list, values)
}

func (list *linkedList[O]) Remove(value O) error {
	node := list.first
	for node != nil {
		if list.converter.Equals(node.value, value) {
			list.size = list.size - 1

			first := node.before != nil
			last := node.after != nil

			if first && last {
				list.first = nil
				list.last = nil
				return nil
			}

			if first {
				list.first = node.after
				node.after.before = nil
				return nil
			}

			if last {
				list.last = node.before
				node.before.after = nil
				return nil
			}

			before := node.before
			after := node.after

			before.after = after
			after.before = before
			return nil
		}

		node = node.after
	}

	return errors.Embed(errors.New(nil, ErrorCodeNotFound, "not found"), value)
}

func (list *linkedList[O]) RemoveAll(values Collection[O]) error {
	return collectionRemoveAll[O](list, values)
}

func (list *linkedList[O]) Size() int {
	return list.size
}

// List implementation

func (list *linkedList[O]) IndexOf(value O) int {
	return listIndexOf[O](list, value, list.converter)
}

func (list *linkedList[O]) Get(ix int) (O, error) {
	var null O
	if ix < 0 || ix >= list.size {
		return null, errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	current := 0
	for iterator := list.Iterator(); iterator.HasNext(); {
		value := iterator.Next()
		if current == ix {
			return value, nil
		}
	}

	return null, errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
}

func (list *linkedList[O]) Insert(value O, ix int) error {
	if ix < 0 || ix > list.size {
		return errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	if list.size == 0 {
		// Then this is empty, and we're adding the first item.
		node := &linkedListNode[O]{
			before: nil,
			after:  nil,
			value:  value,
		}

		list.last = node
		list.first = node
		list.size = 1
		return nil
	}

	if ix == 0 {
		// Then we are inserting at the beginning.
		node := &linkedListNode[O]{
			before: nil,
			after:  list.first,
			value:  value,
		}

		list.first.before = node
		list.first = node
		list.size = list.size + 1
		return nil
	}

	if ix == list.size {
		// Then we are inserting at the end.
		node := &linkedListNode[O]{
			before: list.last,
			after:  nil,
			value:  value,
		}

		list.last.after = node
		list.last = node
		list.size = list.size + 1
		return nil
	}

	var current *linkedListNode[O]

	half := list.size / 2
	if ix < half {
		// Start at the beginning and go forward.
		current = list.first
		currentIx := 0
		for current != nil {
			if currentIx == ix {
				break
			}
			current = current.after
			currentIx = currentIx + 1
		}
	} else {
		// Start at the end and go backward.
		current = list.last
		currentIx := list.size - 1
		for current != nil {
			if currentIx == ix {
				break
			}
			current = current.before
			currentIx = currentIx - 1
		}
	}

	node := &linkedListNode[O]{
		before: current.before,
		after:  current,
		value:  value,
	}

	current.before = node
	list.size = list.size + 1
	return nil
}

func (list *linkedList[O]) Replace(value O, ix int) (O, error) {
	if ix < 0 || ix >= list.size {
		var obj O
		return obj, errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	half := list.size / 2
	var current linkedListNode[O]
	if ix < half {
		count := 0
		for current := list.first; current != nil; current = current.after {
			if count == ix {
			}
		}
	} else {
		count := list.size - 1
		for current := list.last; current != nil; current = current.before {
			if count == ix {
				previous := current.value
				current.value = value
				return previous, nil
			}
		}
	}

	previous := current.value
	current.value = value
	return previous, nil
}

func (list *linkedList[O]) RemoveAt(ix int) error {
	if ix < 0 || ix >= list.size {
		return errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}
}

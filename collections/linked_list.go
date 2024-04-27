package collections

import (
	"encoding/json"
	"sort"

	"github.com/pasataleo/go-errors/errors"
	"github.com/pasataleo/go-objects/objects"
)

type linkedList[O objects.Object] struct {
	first *linkedListNode[O]
	last  *linkedListNode[O]

	size int
}

type linkedListNode[O objects.Object] struct {
	before *linkedListNode[O]
	after  *linkedListNode[O]

	value O
}

func NewLinkedList[O objects.Object]() List[O] {
	return &linkedList[O]{}
}

// Object implementation

func (list *linkedList[O]) Equals(other any) bool {
	return listEquals[O](list, other)
}

func (list *linkedList[O]) HashCode() uint64 {
	return listHashCode[O](list)
}

func (list *linkedList[O]) String() string {
	return listString[O](list)
}

func (list *linkedList[O]) MarshalJSON() ([]byte, error) {
	var values []O
	for iterator := list.Iterator(); iterator.HasNext(); {
		values = append(values, iterator.Next())
	}
	return json.Marshal(values)
}

func (list *linkedList[O]) UnmarshalJSON(bytes []byte) error {
	var values []O
	if err := json.Unmarshal(bytes, &values); err != nil {
		return err
	}
	list.Clear()
	for _, value := range values {
		if err := list.Add(value); err != nil {
			return err
		}
	}
	return nil
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
	node := list.value(value)
	if node == nil {
		return errors.Embed(errors.New(nil, ErrorCodeNotFound, "not found"), value)
	}
	list.remove(node)
	return nil
}

func (list *linkedList[O]) RemoveAll(values Collection[O]) error {
	return collectionRemoveAll[O](list, values)
}

func (list *linkedList[O]) Copy() Collection[O] {
	newList := NewLinkedList[O]()
	for iterator := list.Iterator(); iterator.HasNext(); {
		_ = newList.Add(iterator.Next())
	}
	return newList
}

func (list *linkedList[O]) Size() int {
	return list.size
}

func (list *linkedList[O]) IsEmpty() bool {
	return list.Size() == 0
}

func (list *linkedList[O]) Clear() {
	list.first = nil
	list.last = nil
	list.size = 0
}

// List implementation

func (list *linkedList[O]) IndexOf(value O) int {
	return listIndexOf[O](list, value)
}

func (list *linkedList[O]) Get(ix int) (O, error) {
	node := list.index(ix)
	if node == nil {
		var null O
		return null, errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}
	return node.value, nil
}

func (list *linkedList[O]) Insert(value O, ix int) error {
	if list.size == 0 {
		if ix == 0 {
			node := &linkedListNode[O]{
				value: value,
			}
			list.first = node
			list.last = node
			list.size = 1
			return nil
		}
		return errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	current := list.index(ix)
	if current == nil {
		if ix == list.size {
			// Then we're inserting right at the end.
			node := &linkedListNode[O]{
				value:  value,
				before: list.last,
			}
			list.last.after = node
			list.last = node
			list.size = list.size + 1
			return nil
		}
		return errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	node := &linkedListNode[O]{
		before: current.before,
		after:  current,
		value:  value,
	}

	if current.before == nil {
		// Then this is the first node.
		list.first = node
	}
	current.before = node
	list.size = list.size + 1
	return nil
}

func (list *linkedList[O]) Replace(value O, ix int) (O, error) {
	var previous O

	node := list.index(ix)
	if node == nil {
		return previous, errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	previous = node.value
	node.value = value
	return previous, nil
}

func (list *linkedList[O]) RemoveAt(ix int) (O, error) {
	node := list.index(ix)
	if node == nil {
		var null O
		return null, errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}
	list.remove(node)
	return node.value, nil
}

// linked list

func (list *linkedList[O]) remove(node *linkedListNode[O]) {
	list.size = list.size - 1

	if node.before == nil && node.after == nil {
		// Then there's just a single item in the list, and we're removing it.
		list.first = nil
		list.last = nil
		return
	}

	if node.before == nil {
		// Then we have the first item in the list.
		list.first = node.after
		list.first.before = nil
		return
	}

	if node.after == nil {
		// Then we have the first item in the list.
		list.last = node.before
		list.last.after = nil
		return
	}

	before := node.before
	after := node.after
	before.after = after
	after.before = before
}

func (list *linkedList[O]) value(value O) *linkedListNode[O] {
	for current := list.first; current != nil; current = current.after {
		if current.value.Equals(value) {
			return current
		}
	}
	return nil
}

func (list *linkedList[O]) index(ix int) *linkedListNode[O] {
	mid := list.size / 2
	if ix <= mid {
		for current, currentIx := list.first, 0; current != nil; current, currentIx = current.after, currentIx+1 {
			if currentIx == ix {
				return current
			}
		}
	} else {
		for current, currentIx := list.last, list.size-1; current != nil; current, currentIx = current.before, currentIx-1 {
			if currentIx == ix {
				return current
			}
		}
	}
	return nil
}

func (list *linkedList[O]) sort(comparator objects.Comparator[O]) {
	sorted := make([]O, list.size)
	for iterator := list.Iterator(); iterator.HasNext(); {
		sorted = append(sorted, iterator.Next())
	}
	sort.Slice(sorted, func(i, j int) bool {
		return comparator.Compare(sorted[i], sorted[j]) < 0
	})

	for node, ix := list.first, 0; node != nil; ix++ {
		node.value = sorted[ix]
	}
}

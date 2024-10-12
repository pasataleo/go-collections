package collections

import (
	"encoding/json"
	"iter"
	"sort"

	"github.com/pasataleo/go-errors/errors"
	"github.com/pasataleo/go-objects/objects"
)

type arrayList[O objects.Object] struct {
	values []O
}

// NewArrayList creates a basic array list with the given elements.
func NewArrayList[O objects.Object](elems ...O) List[O] {
	list := &arrayList[O]{}
	list.values = append(list.values, elems...)
	return list
}

// Object implementation

// Equals implements objects.Object.
func (list *arrayList[O]) Equals(other any) bool {
	return listEquals[O](list, other)
}

// HashCode implements objects.Object.
func (list *arrayList[O]) HashCode() uint64 {
	return listHashCode[O](list)
}

// String implements objects.Object.
func (list *arrayList[O]) String() string {
	return listString[O](list)
}

// MarshalJSON implements objects.Object.
func (list *arrayList[O]) MarshalJSON() ([]byte, error) {
	return json.Marshal(list.values)
}

// UnmarshalJSON implements objects.Object.
func (list *arrayList[O]) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &list.values)
}

// Iterable implementation

// Iterator implements objects.Iterable.
func (list *arrayList[O]) Iterator() objects.Iterator[O] {
	return &arrayListIterator[O]{
		current: 0,
		list:    list,
	}
}

// Collection implementation

// Elems implements Collection.
func (list *arrayList[O]) Elems() iter.Seq[O] {
	return objects.SequenceFrom[O](list)
}

// Contains implements Collection.
func (list *arrayList[O]) Contains(value O) bool {
	return list.IndexOf(value) >= 0
}

// ContainsAll implements Collection.
func (list *arrayList[O]) ContainsAll(values Collection[O]) bool {
	return collectionContainsAll[O](list, values)
}

// Add implements Collection.
func (list *arrayList[O]) Add(value O) error {
	// Adding to a list is the same as inserting at the end.
	return list.Insert(value, list.Size())
}

// AddAll implements Collection.
func (list *arrayList[O]) AddAll(values Collection[O]) error {
	return collectionAddAll[O](list, values)
}

// Remove implements Collection.
func (list *arrayList[O]) Remove(value O) error {
	ix := list.IndexOf(value)
	if ix < 0 {
		return errors.Embed(errors.New(nil, ErrorCodeNotFound, "not found"), "value", value)
	}
	_, err := list.RemoveAt(ix)
	return err
}

// RemoveAll implements Collection.
func (list *arrayList[O]) RemoveAll(values Collection[O]) error {
	return collectionRemoveAll[O](list, values)
}

// Copy implements Collection.
func (list *arrayList[O]) Copy() Collection[O] {
	newList := NewArrayList[O]()
	for iterator := list.Iterator(); iterator.HasNext(); {
		_ = newList.Add(iterator.Next())
	}
	return newList
}

// Size implements Collection.
func (list *arrayList[O]) Size() int {
	return len(list.values)
}

// IsEmpty implements Collection.
func (list *arrayList[O]) IsEmpty() bool {
	return list.Size() == 0
}

// Clear implements Collection.
func (list *arrayList[O]) Clear() {
	list.values = nil
}

// List implementation

// IndexOf implements List.
func (list *arrayList[O]) IndexOf(value O) int {
	return listIndexOf[O](list, value)
}

// Get implements List.
func (list *arrayList[O]) Get(ix int) (O, error) {
	if ix < 0 || ix >= len(list.values) {
		var obj O
		return obj, errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	return list.values[ix], nil
}

// Insert implements List.
func (list *arrayList[O]) Insert(value O, ix int) error {
	if ix == len(list.values) {
		list.values = append(list.values, value)
		return nil
	}

	if ix < 0 || ix > len(list.values) {
		return errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	list.values = append(list.values[:ix+1], list.values[ix:]...)
	list.values[ix] = value
	return nil
}

// Replace implements List.
func (list *arrayList[O]) Replace(value O, ix int) (O, error) {
	if ix < 0 || ix >= len(list.values) {
		var obj O
		return obj, errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	current := list.values[ix]
	list.values[ix] = value
	return current, nil
}

// RemoveAt implements List.
func (list *arrayList[O]) RemoveAt(ix int) (O, error) {
	if ix < 0 || ix >= len(list.values) {
		var null O
		return null, errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	obj := list.values[ix]
	list.values = append(list.values[:ix], list.values[ix+1:]...)
	return obj, nil
}

// sort sorts the list using the given comparator.
func (list *arrayList[O]) sort(comparator objects.Comparator[O]) {
	sort.Slice(list.values, func(i, j int) bool {
		return comparator.Compare(list.values[i], list.values[j]) < 0
	})
}

package collections

import (
	"github.com/pasataleo/go-errors/errors"
	"github.com/pasataleo/go-objects/objects"
)

type arrayList[O any] struct {
	values []O

	converter objects.ObjectConverter[O]
}

func NewArrayListT[O any](converter objects.ObjectConverter[O]) List[O] {
	return &arrayList[O]{
		converter: converter,
	}
}

func NewArrayList[O objects.Object]() List[O] {
	return &arrayList[O]{
		converter: objects.ObjectIdentityConverter[O](),
	}
}

// Object implementation

func (list *arrayList[O]) Equals(other any) bool {
	return listEquals[O](list, other, list.converter)
}

func (list *arrayList[O]) HashCode() uint64 {
	return listHashCode[O](list, list.converter)
}

func (list *arrayList[O]) ToString() string {
	return listString[O](list, list.converter)
}

// Iterable implementation

func (list *arrayList[O]) Iterator() objects.Iterator[O] {
	return &arrayListIterator[O]{
		current: 0,
		list:    list,
	}
}

// Collection implementation

func (list *arrayList[O]) Contains(value O) bool {
	return list.IndexOf(value) >= 0
}

func (list *arrayList[O]) ContainsAll(values Collection[O]) bool {
	return collectionContainsAll[O](list, values)
}

func (list *arrayList[O]) Add(value O) error {
	// Adding to a list is the same as inserting at the end.
	return list.Insert(value, list.Size())
}

func (list *arrayList[O]) AddAll(values Collection[O]) error {
	return collectionAddAll[O](list, values)
}

func (list *arrayList[O]) Remove(value O) error {
	ix := list.IndexOf(value)
	if ix < 0 {
		return errors.Embed(errors.New(nil, ErrorCodeNotFound, "not found"), value)
	}
	return list.RemoveAt(ix)
}

func (list *arrayList[O]) RemoveAll(values Collection[O]) error {
	return collectionRemoveAll[O](list, values)
}

func (list *arrayList[O]) Size() int {
	return len(list.values)
}

// List implementation

func (list *arrayList[O]) IndexOf(value O) int {
	return listIndexOf[O](list, value, list.converter)
}

func (list *arrayList[O]) Get(ix int) (O, error) {
	if ix < 0 || ix >= len(list.values) {
		var obj O
		return obj, errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	return list.values[ix], nil
}

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

func (list *arrayList[O]) Replace(value O, ix int) (O, error) {
	if ix < 0 || ix >= len(list.values) {
		var obj O
		return obj, errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	current := list.values[ix]
	list.values[ix] = value
	return current, nil
}

func (list *arrayList[O]) RemoveAt(ix int) error {
	if ix < 0 || ix >= len(list.values) {
		return errors.Newf(nil, ErrorCodeOutOfBounds, "index %d out of bounds", ix)
	}

	list.values = append(list.values[:ix], list.values[ix+1:]...)
	return nil
}

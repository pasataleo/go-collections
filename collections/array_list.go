package collections

import (
	"bytes"
	"fmt"

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
	if lOther, ok := other.(List[O]); ok {
		if lOther.Size() != list.Size() {
			return false
		}

		for ix := 0; ix < list.Size(); ix++ {
			left, err := list.Get(ix)
			if err != nil {
				panic(err)
			}

			right, err := lOther.Get(ix)
			if err != nil {
				panic(err)
			}

			if !list.converter.Equals(left, right) {
				return false
			}
		}
		return true
	}

	return false
}

func (list *arrayList[O]) HashCode() uint64 {
	hash := uint64(13001)
	for ix := 0; ix < list.Size(); ix++ {
		value, err := list.Get(ix)
		if err != nil {
			panic(err)
		}
		hash = hash * list.converter.HashCode(value)
	}
	return hash
}

func (list *arrayList[O]) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for ix := 0; ix < list.Size(); ix++ {
		value, err := list.Get(ix)
		if err != nil {
			// This shouldn't happen, but maybe someone could be editing the
			// list in parallel while we're converting it to a string.
			panic(err)
		}

		if ix == 0 {
			buffer.WriteString(list.converter.ToString(value))
		} else {
			buffer.WriteString(fmt.Sprintf(",%s", list.converter.ToString(value)))
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

// Iterable implementation

func (list *arrayList[O]) Iterator() objects.Iterator[O] {
	return &listIterator[O]{
		current: 0,
		list:    list,
	}
}

// Collection implementation

func (list *arrayList[O]) Contains(value O) bool {
	return list.IndexOf(value) >= 0
}

func (list *arrayList[O]) ContainsAll(values Collection[O]) bool {
	for iterator := values.Iterator(); iterator.HasNext(); {
		value, err := iterator.Next()
		if err != nil {
			// This shouldn't happen as we are checking HasNext first, but
			// something weird could happen with threading.
			panic(err)
		}

		if !list.Contains(value) {
			return false
		}
	}
	return true
}

func (list *arrayList[O]) Add(value O) error {
	// Adding to a list is the same as inserting at the end.
	return list.Insert(value, list.Size())
}

func (list *arrayList[O]) AddAll(values Collection[O]) error {
	var multi error
	for iterator := values.Iterator(); iterator.HasNext(); {
		value, err := iterator.Next()
		if err != nil {
			// This shouldn't happen as we are checking HasNext first, but
			// something weird could happen with threading.
			panic(err)
		}

		if err := list.Add(value); err != nil {
			multi = errors.Append(multi, err)
		}
	}
	return multi
}

func (list *arrayList[O]) Remove(value O) error {
	ix := list.IndexOf(value)
	if ix < 0 {
		return errors.Embed(errors.New(nil, ErrorCodeNotFound, "not found"), value)
	}
	return list.RemoveAt(ix)
}

func (list *arrayList[O]) RemoveAll(values Collection[O]) error {
	var multi error
	for iterator := values.Iterator(); iterator.HasNext(); {
		value, err := iterator.Next()
		if err != nil {
			// This shouldn't really happen unless someone is behaving badly
			// with threads.
			panic(err)
		}

		if err := list.Remove(value); err != nil {
			multi = errors.Append(multi, err)
		}
	}
	return multi
}

func (list *arrayList[O]) Size() int {
	return len(list.values)
}

// List implementation

func (list *arrayList[O]) IndexOf(value O) int {
	for ix := 0; ix < list.Size(); ix++ {
		item, err := list.Get(ix)
		if err != nil {
			// This shouldn't ever happen, but parallelism could be crazy.
			panic(err)
		}

		if list.converter.Equals(item, value) {
			return ix
		}
	}
	return -1
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

package collections

import (
	"bytes"
	"fmt"

	"github.com/pasataleo/go-errors/errors"
	"github.com/pasataleo/go-objects/objects"
)

type hashSet[O any] struct {
	values map[uint64][]O
	size   int

	converter objects.ObjectConverter[O]
}

type hashSetIterator[O any] struct {
	set *hashSet[O]

	keys []uint64
	keyI int

	valueI int
}

func (iterator *hashSetIterator[O]) HasNext() bool {
	return iterator.keyI < len(iterator.keys)
}

func (iterator *hashSetIterator[O]) Next() (O, error) {
	if iterator.keyI < 0 || iterator.keyI >= len(iterator.keys) {
		var obj O
		return obj, errors.New(nil, ErrorCodeOutOfBounds, "out of bounds")
	}

	currentSlice := iterator.set.values[iterator.keys[iterator.keyI]]

	value := currentSlice[iterator.valueI]
	iterator.valueI = iterator.valueI + 1
	if iterator.valueI >= len(currentSlice) {
		iterator.keyI = iterator.keyI + 1
		iterator.valueI = 0
	}
	return value, nil
}

func NewHashSetT[O any](converter objects.ObjectConverter[O]) Set[O] {
	return &hashSet[O]{
		values:    make(map[uint64][]O),
		size:      0,
		converter: converter,
	}
}

func NewHashSet[O objects.Object]() Set[O] {
	return NewHashSetT[O](objects.ObjectIdentityConverter[O]())
}

// Object implementation

func (set *hashSet[O]) Equals(other any) bool {
	if lSet, ok := other.(Set[O]); ok {
		if lSet.Size() != set.Size() {
			return false
		}

		for iterator := set.Iterator(); iterator.HasNext(); {
			value, err := iterator.Next()
			if err != nil {
				panic(err)
			}

			if !lSet.Contains(value) {
				return false
			}
		}
		return true
	}
	return false
}

func (set *hashSet[O]) HashCode() uint64 {
	hash := uint64(13001)
	for iterator := set.Iterator(); iterator.HasNext(); {
		value, err := iterator.Next()
		if err != nil {
			panic(err)
		}

		hash = hash * set.converter.HashCode(value)
	}
	return hash
}

func (set *hashSet[O]) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")

	first := true
	for iterator := set.Iterator(); iterator.HasNext(); {
		value, err := iterator.Next()
		if err != nil {
			panic(err)
		}

		if first {
			buffer.WriteString(set.converter.ToString(value))
		} else {
			buffer.WriteString(fmt.Sprintf(",%s", set.converter.ToString(value)))
		}
		first = false
	}
	buffer.WriteString("]")
	return buffer.String()
}

// Iterable implementation

func (set *hashSet[O]) Iterator() objects.Iterator[O] {
	return &hashSetIterator[O]{
		set: set,
		keys: func() []uint64 {
			var keys []uint64
			for key := range set.values {
				keys = append(keys, key)
			}
			return keys
		}(),
	}
}

// Collection implementation

func (set *hashSet[O]) Contains(value O) bool {
	hash := set.converter.HashCode(value)
	for _, contained := range set.values[hash] {
		if set.converter.Equals(contained, value) {
			return true
		}
	}
	return false
}

func (set *hashSet[O]) ContainsAll(values Collection[O]) bool {
	for iterator := values.Iterator(); iterator.HasNext(); {
		value, err := iterator.Next()
		if err != nil {
			// This shouldn't happen as we are checking HasNext first, but
			// something weird could happen with threading.
			panic(err)
		}

		if !set.Contains(value) {
			return false
		}
	}
	return true
}

func (set *hashSet[O]) Add(value O) error {
	hash := set.converter.HashCode(value)
	values := set.values[hash]
	for _, contained := range values {
		if set.converter.Equals(value, contained) {
			return errors.Embed(errors.New(nil, ErrorCodeAlreadyExists, "already exists"), value)
		}
	}
	values = append(values, value)
	set.values[hash] = values
	set.size = set.size + 1
	return nil
}

func (set *hashSet[O]) AddAll(values Collection[O]) error {
	var multi error
	for iterator := values.Iterator(); iterator.HasNext(); {
		value, err := iterator.Next()
		if err != nil {
			// This shouldn't happen as we are checking HasNext first, but
			// something weird could happen with threading.
			panic(err)
		}

		if err := set.Add(value); err != nil {
			multi = errors.Append(multi, err)
		}
	}
	return multi
}

func (set *hashSet[O]) Remove(value O) error {
	hash := set.converter.HashCode(value)
	values := set.values[hash]
	for ix, contained := range values {
		if set.converter.Equals(value, contained) {
			set.values[hash] = append(values[:ix], values[ix+1:]...)
			set.size = set.size - 1
			return nil
		}
	}
	return errors.Embed(errors.New(nil, ErrorCodeNotFound, "not found"), value)
}

func (set *hashSet[O]) RemoveAll(values Collection[O]) error {
	var multi error
	for iterator := values.Iterator(); iterator.HasNext(); {
		value, err := iterator.Next()
		if err != nil {
			// This shouldn't really happen unless someone is behaving badly
			// with threads.
			panic(err)
		}

		if err := set.Remove(value); err != nil {
			multi = errors.Append(multi, err)
		}
	}
	return multi
}

func (set *hashSet[O]) Size() int {
	return set.size
}

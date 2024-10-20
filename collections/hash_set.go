package collections

import (
	"bytes"
	"encoding/json"
	"fmt"
	"iter"

	"github.com/pasataleo/go-errors/errors"
	"github.com/pasataleo/go-objects/objects"
)

type hashSet[O objects.Object] struct {
	values map[uint64][]O
	size   int
}

// NewHashSet creates a new hash set.
func NewHashSet[O objects.Object]() Set[O] {
	return &hashSet[O]{
		values: make(map[uint64][]O),
		size:   0,
	}
}

// Object implementation

// Equals implements objects.Object.
func (set *hashSet[O]) Equals(other any) bool {
	if lSet, ok := other.(Set[O]); ok {
		if lSet.Size() != set.Size() {
			return false
		}

		for iterator := set.Iterator(); iterator.HasNext(); {
			if !lSet.Contains(iterator.Next()) {
				return false
			}
		}
		return true
	}
	return false
}

// HashCode implements objects.Object.
func (set *hashSet[O]) HashCode() uint64 {
	hash := uint64(13001)
	for iterator := set.Iterator(); iterator.HasNext(); {
		hash = hash * iterator.Next().HashCode()
	}
	return hash
}

// String implements objects.Object.
func (set *hashSet[O]) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("[")

	first := true
	for iterator := set.Iterator(); iterator.HasNext(); {
		value := iterator.Next()
		if first {
			buffer.WriteString(value.String())
		} else {
			buffer.WriteString(fmt.Sprintf(",%s", value))
		}
		first = false
	}
	buffer.WriteString("]")
	return buffer.String()
}

// MarshalJSON implements json.Marshaler.
func (set *hashSet[O]) MarshalJSON() ([]byte, error) {
	var values []O
	for iterator := set.Iterator(); iterator.HasNext(); {
		values = append(values, iterator.Next())
	}
	return json.Marshal(values)
}

// UnmarshalJSON implements json.Unmarshaler.
func (set *hashSet[O]) UnmarshalJSON(bytes []byte) error {
	var values []O
	if err := json.Unmarshal(bytes, &values); err != nil {
		return err
	}

	set.Clear()
	for _, value := range values {
		if err := set.Add(value); err != nil {
			return err
		}
	}
	return nil
}

// Iterable implementation

// Iterator implements objects.Iterable.
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

// Elems implements Collection.
func (set *hashSet[O]) Elems() iter.Seq[O] {
	return func(yield func(O) bool) {
		for _, values := range set.values {
			for _, value := range values {
				yield(value)
			}
		}
	}
}

// Contains implements Collection.
func (set *hashSet[O]) Contains(value O) bool {
	hash := value.HashCode()
	for _, contained := range set.values[hash] {
		if contained.Equals(value) {
			return true
		}
	}
	return false
}

// ContainsAll implements Collection.
func (set *hashSet[O]) ContainsAll(values Collection[O]) bool {
	return collectionContainsAll[O](set, values)
}

// Add implements Collection.
func (set *hashSet[O]) Add(value O) error {
	hash := value.HashCode()
	values := set.values[hash]
	for _, contained := range values {
		if value.Equals(contained) {
			return errors.Embed(errors.New(nil, ErrorCodeAlreadyExists, "already exists"), "value", value)
		}
	}
	values = append(values, value)
	set.values[hash] = values
	set.size = set.size + 1
	return nil
}

// AddAll implements Collection.
func (set *hashSet[O]) AddAll(values Collection[O]) error {
	return collectionAddAll[O](set, values)
}

// Remove implements Collection.
func (set *hashSet[O]) Remove(value O) error {
	hash := value.HashCode()
	values := set.values[hash]
	for ix, contained := range values {
		if value.Equals(contained) {
			set.values[hash] = append(values[:ix], values[ix+1:]...)
			set.size = set.size - 1
			return nil
		}
	}
	return errors.Embed(errors.New(nil, ErrorCodeNotFound, "not found"), "value", value)
}

// RemoveAll implements Collection.
func (set *hashSet[O]) RemoveAll(values Collection[O]) error {
	return collectionRemoveAll[O](set, values)
}

// Copy implements Collection.
func (set *hashSet[O]) Copy() Collection[O] {
	newSet := NewHashSet[O]()
	for iterator := set.Iterator(); iterator.HasNext(); {
		_ = newSet.Add(iterator.Next())
	}
	return newSet
}

// Size implements Collection.
func (set *hashSet[O]) Size() int {
	return set.size
}

// IsEmpty implements Collection.
func (set *hashSet[O]) IsEmpty() bool {
	return set.Size() == 0
}

// Clear implements Collection.
func (set *hashSet[O]) Clear() {
	set.values = make(map[uint64][]O)
	set.size = 0
}

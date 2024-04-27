package collections

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pasataleo/go-errors/errors"
	"github.com/pasataleo/go-objects/objects"
)

type hashMap[K objects.Object, V objects.Object] struct {
	values map[uint64][]MapEntry[K, V]
	size   int
}

func NewHashMap[K objects.Object, V objects.Object]() Map[K, V] {
	return &hashMap[K, V]{
		values: make(map[uint64][]MapEntry[K, V]),
	}
}

// Object implementation

func (h *hashMap[K, V]) Equals(other any) bool {
	if lMap, ok := other.(Map[K, V]); ok {
		if lMap.Size() != h.Size() {
			return false
		}

		for iterator := h.Iterator(); iterator.HasNext(); {
			entry := iterator.Next()

			if !h.ContainsKey(entry.GetKey()) {
				return false
			}

			if !entry.GetValue().Equals(h.Get(entry.GetKey())) {
				return false
			}
		}
		return true
	}
	return false
}

func (h *hashMap[K, V]) HashCode() uint64 {
	hash := uint64(13001)
	for iterator := h.Iterator(); iterator.HasNext(); {
		entry := iterator.Next()
		hash = hash * entry.HashCode()
	}
	return hash
}

func (h *hashMap[K, V]) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{")

	first := true
	for iterator := h.Iterator(); iterator.HasNext(); {
		entry := iterator.Next()

		if first {
			buffer.WriteString(entry.String())
		} else {
			buffer.WriteString(fmt.Sprintf(",%s", entry))
		}
		first = false
	}
	buffer.WriteString("}")
	return buffer.String()
}

func (h *hashMap[K, V]) MarshalJSON() ([]byte, error) {
	var entries []MapEntry[K, V]
	for iterator := h.Iterator(); iterator.HasNext(); {
		entries = append(entries, iterator.Next())
	}
	return json.Marshal(entries)
}

func (h *hashMap[K, V]) UnmarshalJSON(bytes []byte) error {
	var entries []MapEntry[K, V]
	if err := json.Unmarshal(bytes, &entries); err != nil {
		return err
	}

	h.Clear()
	for _, entry := range entries {
		if err := h.Put(entry.GetKey(), entry.GetValue()); err != nil {
			return err
		}
	}
	return nil
}

// Iterable implementation

func (h *hashMap[K, V]) Iterator() objects.Iterator[MapEntry[K, V]] {
	return &hashMapIterator[K, V]{
		set: h,
		keys: func() []uint64 {
			var keys []uint64
			for key := range h.values {
				keys = append(keys, key)
			}
			return keys
		}(),
	}
}

// Collection implementation

func (h *hashMap[K, V]) Add(value MapEntry[K, V]) error {
	return h.Put(value.GetKey(), value.GetValue())
}

func (h *hashMap[K, V]) AddAll(values Collection[MapEntry[K, V]]) error {
	return collectionAddAll[MapEntry[K, V]](h, values)
}

func (h *hashMap[K, V]) Remove(value MapEntry[K, V]) error {
	if !h.Contains(value) {
		return errors.Embed(errors.New(nil, ErrorCodeNotFound, "not found"), value)
	}

	_, err := h.Delete(value.GetKey())
	return err
}

func (h *hashMap[K, V]) RemoveAll(values Collection[MapEntry[K, V]]) error {
	return collectionRemoveAll[MapEntry[K, V]](h, values)
}

func (h *hashMap[K, V]) Contains(value MapEntry[K, V]) bool {
	if !h.ContainsKey(value.GetKey()) {
		return false
	}

	contained := h.Get(value.GetKey())
	return value.GetValue().Equals(contained)
}

func (h *hashMap[K, V]) ContainsAll(values Collection[MapEntry[K, V]]) bool {
	return collectionContainsAll[MapEntry[K, V]](h, values)
}

func (h *hashMap[K, V]) Copy() Collection[MapEntry[K, V]] {
	newMap := NewHashMap[K, V]()
	for iterator := h.Iterator(); iterator.HasNext(); {
		_ = newMap.Add(iterator.Next())
	}
	return newMap
}

func (h *hashMap[K, V]) Size() int {
	return h.size
}

func (h *hashMap[K, V]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *hashMap[K, V]) Clear() {
	h.values = make(map[uint64][]MapEntry[K, V])
	h.size = 0
}

// Map implementation

func (h *hashMap[K, V]) ContainsKey(key K) bool {
	hash := key.HashCode()

	values := h.values[hash]
	for _, entry := range values {
		if key.Equals(entry.GetKey()) {
			return true
		}
	}
	return false
}

func (h *hashMap[K, V]) Put(key K, value V) error {
	hash := key.HashCode()

	newEntry := &mapEntry[K, V]{
		Key:   key,
		Value: value,
	}

	values := h.values[hash]
	for _, entry := range values {
		if key.Equals(entry.GetKey()) {
			return errors.Embed(errors.New(nil, ErrorCodeAlreadyExists, "already exists"), key)
		}
	}
	values = append(values, newEntry)
	h.values[hash] = values
	h.size = h.size + 1
	return nil
}

func (h *hashMap[K, V]) Replace(key K, value V) (V, error) {
	hash := key.HashCode()

	newEntry := &mapEntry[K, V]{
		Key:   key,
		Value: value,
	}

	values := h.values[hash]
	for ix, entry := range values {
		if key.Equals(entry.GetKey()) {
			oldValue := entry.GetValue()
			values[ix] = newEntry
			h.values[hash] = values
			return oldValue, nil
		}
	}

	var obj V
	return obj, errors.Embed(errors.New(nil, ErrorCodeNotFound, "not found"), key)
}

func (h *hashMap[K, V]) PutOrReplace(key K, value V) (V, bool) {
	hash := key.HashCode()

	newEntry := &mapEntry[K, V]{
		Key:   key,
		Value: value,
	}

	values := h.values[hash]
	for ix, entry := range values {
		if key.Equals(key) {
			oldValue := entry.GetValue()
			values[ix] = newEntry
			h.values[hash] = values
			return oldValue, true
		}
	}

	values = append(values, newEntry)
	h.values[hash] = values
	h.size = h.size + 1
	return value, false
}

func (h *hashMap[K, V]) Delete(key K) (V, error) {
	hash := key.HashCode()

	values := h.values[hash]
	for ix, entry := range values {
		if key.Equals(entry.GetKey()) {
			newValues := append(values[:ix], values[ix+1:]...)
			h.values[hash] = newValues
			h.size = h.size - 1
			return entry.GetValue(), nil
		}
	}

	var obj V
	return obj, errors.Embed(errors.New(nil, ErrorCodeNotFound, "not found"), key)
}

func (h *hashMap[K, V]) DeleteIfPresent(key K) (V, bool) {
	hash := key.HashCode()

	values := h.values[hash]
	for ix, entry := range values {
		if key.Equals(entry.GetKey()) {
			newValues := append(values[:ix], values[ix+1:]...)
			h.values[hash] = newValues
			h.size = h.size - 1
			return entry.GetValue(), true
		}
	}

	var obj V
	return obj, false
}

func (h *hashMap[K, V]) Get(key K) V {
	hash := key.HashCode()

	values := h.values[hash]
	for _, entry := range values {
		if key.Equals(entry.GetKey()) {
			return entry.GetValue()
		}
	}
	panic("not found")
}

func (h *hashMap[K, V]) GetSafe(key K) (V, error) {
	hash := key.HashCode()

	values := h.values[hash]
	for _, entry := range values {
		if key.Equals(entry.GetKey()) {
			return entry.GetValue(), nil
		}
	}

	var obj V
	return obj, errors.Embed(errors.New(nil, ErrorCodeNotFound, "not found"), key)
}

func (h *hashMap[K, V]) Keys() Collection[K] {
	set := NewHashSet[K]()
	for iterator := h.Iterator(); iterator.HasNext(); {
		value := iterator.Next()
		if err := set.Add(value.GetKey()); err != nil {
			panic(err)
		}
	}
	return set
}

func (h *hashMap[K, V]) Values() Collection[V] {
	list := NewArrayList[V]()
	for iterator := h.Iterator(); iterator.HasNext(); {
		value := iterator.Next()
		if err := list.Add(value.GetValue()); err != nil {
			panic(err)
		}
	}
	return list
}

package collections

import (
	"encoding/json"
	"fmt"
	"iter"

	"github.com/pasataleo/go-objects/objects"
)

// MapEntry represents a key-value pair in a map.
type MapEntry[K, V objects.Object] interface {
	objects.Object

	GetKey() K
	GetValue() V
}

// Map represents a collection of key-value pairs.
type Map[K, V objects.Object] interface {
	Collection[MapEntry[K, V]]

	// Entries returns an iterator over the entries in the map.
	Entries() iter.Seq2[K, V]

	// ContainsKey returns true if the map contains the given key.
	ContainsKey(key K) bool

	// Put inserts a key-value pair into the map.
	Put(key K, value V) error

	// Replace replaces the value associated with the given key.
	Replace(key K, value V) (V, error)

	// PutOrReplace inserts or replaces a key-value pair in the map.
	PutOrReplace(key K, value V) (V, bool)

	// Delete removes the key-value pair associated with the given key.
	Delete(key K) (V, error)

	// DeleteIfPresent removes the key-value pair associated with the given key if it exists.
	DeleteIfPresent(key K) (V, bool)

	// Get returns the value associated with the given key.
	Get(key K) V

	// GetSafe returns the value associated with the given key or an error if the key does not exist.
	GetSafe(key K) (V, error)

	// Keys returns a collection of the keys in the map.
	Keys() Collection[K]

	// Values returns a collection of the values in the map.
	Values() Collection[V]
}

type mapEntry[K, V objects.Object] struct {
	Key   K `json:"key"`
	Value V `json:"value"`
}

func (m *mapEntry[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m)
}

func (m *mapEntry[K, V]) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, m)
}

func (m *mapEntry[K, V]) Equals(other any) bool {
	if mOther, ok := other.(MapEntry[K, V]); ok {
		return m.Key.Equals(mOther.GetKey()) && m.Value.Equals(mOther.GetValue())
	}
	return false
}

func (m *mapEntry[K, V]) HashCode() uint64 {
	return 37 * m.Key.HashCode() * m.Value.HashCode()
}

func (m *mapEntry[K, V]) String() string {
	return fmt.Sprintf("%s:%s", m.Key, m.Value)
}

func (m *mapEntry[K, V]) GetKey() K {
	return m.Key
}

func (m *mapEntry[K, V]) GetValue() V {
	return m.Value
}

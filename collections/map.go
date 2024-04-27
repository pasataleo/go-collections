package collections

import (
	"encoding/json"
	"fmt"

	"github.com/pasataleo/go-objects/objects"
)

type MapEntry[K, V objects.Object] interface {
	objects.Object

	GetKey() K
	GetValue() V
}

type Map[K, V objects.Object] interface {
	Collection[MapEntry[K, V]]

	ContainsKey(key K) bool

	Put(key K, value V) error
	Replace(key K, value V) (V, error)
	PutOrReplace(key K, value V) (V, bool)

	Delete(key K) (V, error)
	DeleteIfPresent(key K) (V, bool)

	Get(key K) V
	GetSafe(key K) (V, error)

	Keys() Collection[K]
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

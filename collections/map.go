package collections

import (
	"fmt"

	"github.com/pasataleo/go-objects/objects"
)

type MapEntry[K, V any] interface {
	objects.Object

	GetKey() K
	GetValue() V
}

type Map[K, V any] interface {
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

type mapEntry[K, V any] struct {
	key   K
	value V

	keyConverter   objects.ObjectConverter[K]
	valueConverter objects.ObjectConverter[V]
}

func (m mapEntry[K, V]) Equals(other any) bool {
	if mOther, ok := other.(MapEntry[K, V]); ok {
		return m.keyConverter.Equals(m.key, mOther.GetKey()) && m.valueConverter.Equals(m.value, mOther.GetValue())
	}
	return false
}

func (m mapEntry[K, V]) HashCode() uint64 {
	return 37 * m.keyConverter.HashCode(m.key) * m.valueConverter.HashCode(m.value)
}

func (m mapEntry[K, V]) ToString() string {
	return fmt.Sprintf("%s:%s", m.keyConverter.ToString(m.key), m.valueConverter.ToString(m.value))
}

func (m mapEntry[K, V]) GetKey() K {
	return m.key
}

func (m mapEntry[K, V]) GetValue() V {
	return m.value
}

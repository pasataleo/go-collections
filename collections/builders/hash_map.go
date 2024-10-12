package builders

import (
	"github.com/pasataleo/go-objects/objects"

	"github.com/pasataleo/go-collections/collections"
)

type HashMapBuilder[K, V objects.Object] struct {
	collection collections.Map[K, V]
}

func NewHashMapBuilder[K, V objects.Object]() *HashMapBuilder[K, V] {
	return &HashMapBuilder[K, V]{
		collection: collections.NewHashMap[K, V](),
	}
}

func (b *HashMapBuilder[K, V]) Put(key K, value V) *HashMapBuilder[K, V] {
	b.collection.Put(key, value)
	return b
}

func (b *HashMapBuilder[K, V]) Build() collections.Map[K, V] {
	return b.collection
}

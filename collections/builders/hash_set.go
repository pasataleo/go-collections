package builders

import (
	"github.com/pasataleo/go-objects/objects"

	"github.com/pasataleo/go-collections/collections"
)

type HashSetBuilder[V objects.Object] struct {
	collection collections.Set[V]
}

func NewHashSetBuilder[V objects.Object]() *HashSetBuilder[V] {
	return &HashSetBuilder[V]{
		collection: collections.NewHashSet[V](),
	}
}

func (b *HashSetBuilder[V]) Add(value V) *HashSetBuilder[V] {
	b.collection.Add(value)
	return b
}

func (b *HashSetBuilder[V]) Build() collections.Set[V] {
	return b.collection
}

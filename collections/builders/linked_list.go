package builders

import (
	"github.com/pasataleo/go-objects/objects"

	"github.com/pasataleo/go-collections/collections"
)

type LinkedListBuilder[V objects.Object] struct {
	collection collections.List[V]
}

func NewLinkedListBuilder[V objects.Object]() *LinkedListBuilder[V] {
	return &LinkedListBuilder[V]{
		collection: collections.NewLinkedList[V](),
	}
}

func (b *LinkedListBuilder[V]) Add(value V) *LinkedListBuilder[V] {
	b.collection.Add(value)
	return b
}

func (b *LinkedListBuilder[V]) Build() collections.List[V] {
	return b.collection
}

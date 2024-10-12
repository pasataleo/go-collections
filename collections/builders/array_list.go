package builders

import (
	"github.com/pasataleo/go-objects/objects"

	"github.com/pasataleo/go-collections/collections"
)

type ArrayListBuilder[V objects.Object] struct {
	collection collections.List[V]
}

func NewArrayListBuilder[V objects.Object]() *ArrayListBuilder[V] {
	return &ArrayListBuilder[V]{
		collection: collections.NewArrayList[V](),
	}
}

func (builder *ArrayListBuilder[V]) Add(value V) *ArrayListBuilder[V] {
	builder.collection.Add(value)
	return builder
}

func (builder *ArrayListBuilder[V]) Build() collections.List[V] {
	return builder.collection
}

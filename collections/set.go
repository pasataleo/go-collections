package collections

import "github.com/pasataleo/go-objects/objects"

// Set is a collection of unique objects.
type Set[O objects.Object] interface {
	Collection[O]
}

package collections

import "github.com/pasataleo/go-objects/objects"

type Set[O objects.Object] interface {
	Collection[O]
}

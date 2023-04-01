package collections

import (
	"github.com/pasataleo/go-objects/objects"
)

type Collection[O any] interface {
	objects.Object
	objects.Iterable[O]

	Add(value O) error
	AddAll(values Collection[O]) error

	Remove(value O) error
	RemoveAll(values Collection[O]) error

	Contains(value O) bool
	ContainsAll(value Collection[O]) bool

	Size() int
}

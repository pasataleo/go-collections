package collections

import "github.com/pasataleo/go-objects/objects"

type Stack[O objects.Object] interface {
	Collection[O]

	Offer(value O) error
	Peep() (O, error)
	Pop() (O, error)
}

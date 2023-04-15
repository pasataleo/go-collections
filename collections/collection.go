package collections

import (
	"github.com/pasataleo/go-errors/errors"
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
	ContainsAll(values Collection[O]) bool

	Copy() Collection[O]

	Size() int
	IsEmpty() bool
}

func collectionContainsAll[O any](collection Collection[O], target Collection[O]) bool {
	for iterator := target.Iterator(); iterator.HasNext(); {
		if !collection.Contains(iterator.Next()) {
			return false
		}
	}
	return true
}

func collectionAddAll[O any](collection Collection[O], target Collection[O]) error {
	var multi error
	for iterator := target.Iterator(); iterator.HasNext(); {
		if err := collection.Add(iterator.Next()); err != nil {
			multi = errors.Append(multi, err)
		}
	}
	return multi
}

func collectionRemoveAll[O any](collection Collection[O], target Collection[O]) error {
	var multi error
	for iterator := target.Iterator(); iterator.HasNext(); {
		if err := collection.Remove(iterator.Next()); err != nil {
			multi = errors.Append(multi, err)
		}
	}
	return multi
}

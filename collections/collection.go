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
	ContainsAll(value Collection[O]) bool

	Size() int
}

func collectionContainsAll[O any](collection Collection[O], target Collection[O]) bool {
	for iterator := target.Iterator(); iterator.HasNext(); {
		value, err := iterator.Next()
		if err != nil {
			// This shouldn't happen as we are checking HasNext first, but
			// something weird could happen with threading.
			panic(err)
		}

		if !collection.Contains(value) {
			return false
		}
	}
	return true
}

func collectionAddAll[O any](collection Collection[O], target Collection[O]) error {
	var multi error
	for iterator := target.Iterator(); iterator.HasNext(); {
		value, err := iterator.Next()
		if err != nil {
			// This shouldn't happen as we are checking HasNext first, but
			// something weird could happen with threading.
			panic(err)
		}

		if err := collection.Add(value); err != nil {
			multi = errors.Append(multi, err)
		}
	}
	return multi
}

func collectionRemoveAll[O any](collection Collection[O], target Collection[O]) error {
	var multi error
	for iterator := target.Iterator(); iterator.HasNext(); {
		value, err := iterator.Next()
		if err != nil {
			// This shouldn't really happen unless someone is behaving badly
			// with threads.
			panic(err)
		}

		if err := collection.Remove(value); err != nil {
			multi = errors.Append(multi, err)
		}
	}
	return multi
}

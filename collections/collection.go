package collections

import (
	"iter"

	"github.com/pasataleo/go-errors/errors"
	"github.com/pasataleo/go-objects/objects"
)

// Collection is a generic collection of objects.
type Collection[O objects.Object] interface {
	objects.Object
	objects.Iterable[O]

	// Elems returns a sequence of all elements in the collection.
	Elems() iter.Seq[O]

	// Add adds the given value to the collection, and returns an error if the element couldn't be added.
	Add(value O) error

	// AddAll adds all the values in the given collection to the collection, and returns an error if any element
	// couldn't be added.
	AddAll(values Collection[O]) error

	// Remove removes the given value from the collection, and returns an error if the element couldn't be removed.
	Remove(value O) error

	// RemoveAll removes all the values in the given collection from the collection, and returns an error if any element
	// couldn't be removed.
	RemoveAll(values Collection[O]) error

	// Contains returns true if the collection contains the given value.
	Contains(value O) bool

	// ContainsAll returns true if the collection contains all the values in the given collection.
	ContainsAll(values Collection[O]) bool

	// Copy returns a copy of the collection. This should return the same implementation type as the original collection.
	Copy() Collection[O]

	// Size returns the number of elements in the collection.
	Size() int

	// IsEmpty returns true if the collection is empty.
	IsEmpty() bool

	// Clear removes all elements from the collection.
	Clear()
}

func collectionContainsAll[O objects.Object](collection Collection[O], target Collection[O]) bool {
	for iterator := target.Iterator(); iterator.HasNext(); {
		if !collection.Contains(iterator.Next()) {
			return false
		}
	}
	return true
}

func collectionAddAll[O objects.Object](collection Collection[O], target Collection[O]) error {
	var multi error
	for iterator := target.Iterator(); iterator.HasNext(); {
		if err := collection.Add(iterator.Next()); err != nil {
			multi = errors.Append(multi, err)
		}
	}
	return multi
}

func collectionRemoveAll[O objects.Object](collection Collection[O], target Collection[O]) error {
	var multi error
	for iterator := target.Iterator(); iterator.HasNext(); {
		if err := collection.Remove(iterator.Next()); err != nil {
			multi = errors.Append(multi, err)
		}
	}
	return multi
}

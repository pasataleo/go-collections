package collections

import "github.com/pasataleo/go-objects/objects"

// hashSetIterator is an iterator for hashSet.
type hashSetIterator[O objects.Object] struct {
	set *hashSet[O]

	keys []uint64
	keyI int

	valueI int
}

// HasNext implements objects.Iterator.
func (iterator *hashSetIterator[O]) HasNext() bool {
	return iterator.keyI < len(iterator.keys)
}

// Next implements objects.Iterator.
func (iterator *hashSetIterator[O]) Next() O {
	if iterator.keyI < 0 || iterator.keyI >= len(iterator.keys) {
		panic("out of bounds")
	}

	currentSlice := iterator.set.values[iterator.keys[iterator.keyI]]

	value := currentSlice[iterator.valueI]
	iterator.valueI = iterator.valueI + 1
	if iterator.valueI >= len(currentSlice) {
		iterator.keyI = iterator.keyI + 1
		iterator.valueI = 0
	}
	return value
}

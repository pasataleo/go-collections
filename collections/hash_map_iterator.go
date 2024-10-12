package collections

import "github.com/pasataleo/go-objects/objects"

type hashMapIterator[K, V objects.Object] struct {
	set *hashMap[K, V]

	keys []uint64
	keyI int

	valueI int
}

// HasNext implements objects.Iterator.
func (iterator *hashMapIterator[K, V]) HasNext() bool {
	return iterator.keyI < len(iterator.keys)
}

// Next implements objects.Iterator.
func (iterator *hashMapIterator[K, V]) Next() MapEntry[K, V] {
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

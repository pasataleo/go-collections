package collections

import "github.com/pasataleo/go-objects/objects"

type Queue[O any] interface {
	Collection[O]

	Offer(value O) error
	Peep() (O, error)
	Pop() (O, error)
}

func queueEquals[O any](left, right any, converter objects.ObjectConverter[O]) bool {
	lQueue, lOk := left.(Queue[O])
	rQueue, rOk := right.(Queue[O])

	if !lOk || !rOk {
		return false
	}

	if lQueue.Size() != rQueue.Size() {
		return false
	}

	lIterator := lQueue.Iterator()
	rIterator := rQueue.Iterator()

	for lIterator.HasNext() && rIterator.HasNext() {
		l := lIterator.Next()
		r := rIterator.Next()

		if !converter.Equals(l, r) {
			return false
		}
	}
	return lIterator.HasNext() == rIterator.HasNext()
}

func queueHashCode[O any](queue Queue[O], converter objects.ObjectConverter[O]) uint64 {
	hashcode := uint64(13997)
	for iterator := queue.Iterator(); iterator.HasNext(); {
		hashcode = hashcode * converter.HashCode(iterator.Next())
	}
	return hashcode
}

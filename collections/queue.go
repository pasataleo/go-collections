package collections

import "github.com/pasataleo/go-objects/objects"

type Queue[O objects.Object] interface {
	Collection[O]

	Offer(value O) error
	Peep() (O, error)
	Pop() (O, error)
}

func queueEquals[O objects.Object](left, right any) bool {
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

		if !l.Equals(r) {
			return false
		}
	}
	return lIterator.HasNext() == rIterator.HasNext()
}

func queueHashCode[O objects.Object](queue Queue[O]) uint64 {
	hashcode := uint64(13997)
	for iterator := queue.Iterator(); iterator.HasNext(); {
		hashcode = hashcode * iterator.Next().HashCode()
	}
	return hashcode
}

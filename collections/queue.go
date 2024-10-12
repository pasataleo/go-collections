package collections

import "github.com/pasataleo/go-objects/objects"

// Queue is a collection that orders elements in a FIFO (first-in-first-out) manner.
type Queue[O objects.Object] interface {
	Collection[O]

	// Offer adds an element to the end of the queue.
	Offer(value O) error

	// Peep returns the element at the front of the queue without removing it.
	Peep() (O, error)

	// Pop removes and returns the element at the front of the queue.
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

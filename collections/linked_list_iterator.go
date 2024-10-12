package collections

import "github.com/pasataleo/go-objects/objects"

// linkedListIterator is an iterator for a linked list.
type linkedListIterator[O objects.Object] struct {
	current *linkedListNode[O]
}

// HasNext implements objects.Iterator.
func (iterator *linkedListIterator[O]) HasNext() bool {
	return iterator.current != nil
}

// Next implements objects.Iterator.
func (iterator *linkedListIterator[O]) Next() O {
	if iterator.current == nil {
		panic("out of bounds")
	}
	current := iterator.current
	iterator.current = current.after
	return current.value
}

package collections

type linkedListIterator[O any] struct {
	current *linkedListNode[O]
}

func (iterator *linkedListIterator[O]) HasNext() bool {
	return iterator.current != nil
}

func (iterator *linkedListIterator[O]) Next() O {
	if iterator.current == nil {
		panic("out of bounds")
	}
	current := iterator.current
	iterator.current = current.after
	return current.value
}

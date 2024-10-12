package collections

import "github.com/pasataleo/go-objects/objects"

// arrayListIterator is an iterator for the ArrayList.
type arrayListIterator[O objects.Object] struct {
	current int
	list    *arrayList[O]
}

// HasNext implements objects.Iterator.
func (iterator *arrayListIterator[O]) HasNext() bool {
	return iterator.current < iterator.list.Size()
}

// Next implements objects.Iterator.
func (iterator *arrayListIterator[O]) Next() O {
	item, err := iterator.list.Get(iterator.current)
	if err != nil {
		panic(err)
	}
	iterator.current = iterator.current + 1
	return item
}

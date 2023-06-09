package collections

type arrayListIterator[O any] struct {
	current int
	list    *arrayList[O]
}

func (iterator *arrayListIterator[O]) HasNext() bool {
	return iterator.current < iterator.list.Size()
}

func (iterator *arrayListIterator[O]) Next() O {
	item, err := iterator.list.Get(iterator.current)
	if err != nil {
		panic(err)
	}
	iterator.current = iterator.current + 1
	return item
}

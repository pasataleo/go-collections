package collections

type listIterator[O any] struct {
	current int
	list    List[O]
}

func (iterator *listIterator[O]) HasNext() bool {
	return iterator.current < iterator.list.Size()
}

func (iterator *listIterator[O]) Next() (O, error) {
	item, err := iterator.list.Get(iterator.current)
	if err != nil {
		return item, err
	}
	iterator.current = iterator.current + 1
	return item, nil
}

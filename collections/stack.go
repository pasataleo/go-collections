package collections

type Stack[O any] interface {
	Collection[O]

	Offer(value O) error
	Peep() (O, error)
	Pop() (O, error)
}

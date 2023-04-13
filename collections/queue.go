package collections

type Queue[O any] interface {
	Offer(value O) error
	Poll() (O, error)
	Pop() (O, error)
}

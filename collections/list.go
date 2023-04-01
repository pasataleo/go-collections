package collections

type List[O any] interface {
	Collection[O]

	IndexOf(value O) int

	Get(ix int) (O, error)

	Insert(value O, ix int) error
	Replace(value O, ix int) (O, error)
	RemoveAt(ix int) error
}

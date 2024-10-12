package collections

import "github.com/pasataleo/go-objects/objects"

// Stack is a collection that follows the LIFO (Last In First Out) principle.
type Stack[O objects.Object] interface {
	Collection[O]

	// Offer pushes a value onto the stack.
	Offer(value O) error

	// Peep returns the value at the top of the stack without removing it.
	Peep() (O, error)

	// Pop removes and returns the value at the top of the stack.
	Pop() (O, error)
}

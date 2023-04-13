package collections

import (
	"bytes"
	"fmt"

	"github.com/pasataleo/go-objects/objects"
)

type List[O any] interface {
	Collection[O]

	IndexOf(value O) int

	Get(ix int) (O, error)

	Insert(value O, ix int) error
	Replace(value O, ix int) (O, error)
	RemoveAt(ix int) error
}

func listEquals[O any](target List[O], right any, converter objects.ObjectConverter[O]) bool {
	other, ok := right.(List[O])
	if !ok {
		return false
	}

	if target.Size() != other.Size() {
		return false
	}

	var l, r objects.Iterator[O]
	for l, r = target.Iterator(), other.Iterator(); l.HasNext() && r.HasNext(); {
		if !converter.Equals(l.Next(), r.Next()) {
			return false
		}
	}
	return true
}

func listHashCode[O any](list List[O], converter objects.ObjectConverter[O]) uint64 {
	hash := uint64(13001)
	for ix := 0; ix < list.Size(); ix++ {
		value, err := list.Get(ix)
		if err != nil {
			panic(err)
		}
		hash = hash * converter.HashCode(value)
	}
	return hash
}

func listString[O any](list List[O], converter objects.ObjectConverter[O]) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for ix := 0; ix < list.Size(); ix++ {
		value, err := list.Get(ix)
		if err != nil {
			// This shouldn't happen, but maybe someone could be editing the
			// list in parallel while we're converting it to a string.
			panic(err)
		}

		if ix == 0 {
			buffer.WriteString(converter.ToString(value))
		} else {
			buffer.WriteString(fmt.Sprintf(",%s", converter.ToString(value)))
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

func listIndexOf[O any](list List[O], value O, converter objects.ObjectConverter[O]) int {
	ix := 0
	for iterator := list.Iterator(); iterator.HasNext(); {
		if converter.Equals(iterator.Next(), value) {
			return ix
		}
		ix++
	}
	return -1
}

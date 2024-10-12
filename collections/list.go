package collections

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/pasataleo/go-objects/objects"
)

// List represents a collection of objects that can be accessed by index.
type List[O objects.Object] interface {
	Collection[O]

	// IndexOf returns the index of the first occurrence of the given value in the list, or -1 if the value isn't in
	// the list.
	IndexOf(value O) int

	// Get returns the value at the given index.
	Get(ix int) (O, error)

	// Insert adds the given value at the given index.
	Insert(value O, ix int) error

	// Replace replaces the value at the given index with the given value.
	Replace(value O, ix int) (O, error)

	// RemoveAt removes the value at the given index.
	RemoveAt(ix int) (O, error)
}

func listEquals[O objects.Object](target List[O], right any) bool {
	other, ok := right.(List[O])
	if !ok {
		return false
	}

	if target.Size() != other.Size() {
		return false
	}

	var l, r objects.Iterator[O]
	for l, r = target.Iterator(), other.Iterator(); l.HasNext() && r.HasNext(); {
		if !l.Next().Equals(r.Next()) {
			return false
		}
	}
	return true
}

func listHashCode[O objects.Object](list List[O]) uint64 {
	hash := uint64(13001)
	for ix := 0; ix < list.Size(); ix++ {
		value, err := list.Get(ix)
		if err != nil {
			panic(err)
		}
		hash = hash * value.HashCode()
	}
	return hash
}

func listString[O objects.Object](list List[O]) string {
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
			buffer.WriteString(value.String())
		} else {
			buffer.WriteString(fmt.Sprintf(",%s", value))
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

func listIndexOf[O objects.Object](list List[O], value O) int {
	ix := 0
	for iterator := list.Iterator(); iterator.HasNext(); {
		if iterator.Next().Equals(value) {
			return ix
		}
		ix++
	}
	return -1
}

// Sort sorts the given list in ascending order.
func Sort[O objects.ComparableObject[O]](list List[O]) {
	SortT(list, objects.ComparableComparator[O]())
}

// SortT sorts the given list in ascending order using the given comparator.
func SortT[O objects.Object](list List[O], comparator objects.Comparator[O]) {
	switch l := list.(type) {
	case *linkedList[O]:
		l.sort(comparator)
		return
	case *arrayList[O]:
		l.sort(comparator)
		return
	}

	sorted := make([]O, list.Size())
	for iterator := list.Iterator(); iterator.HasNext(); {
		sorted = append(sorted, iterator.Next())
	}
	sort.Slice(sorted, func(i, j int) bool {
		return comparator.Compare(sorted[i], sorted[j]) < 0
	})

	for ix := 0; ix < list.Size(); ix++ {
		if _, err := list.Replace(sorted[ix], ix); err != nil {
			panic(err)
		}
	}
}

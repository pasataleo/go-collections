package collections

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/pasataleo/go-objects/objects"
)

type List[O objects.Object] interface {
	Collection[O]

	IndexOf(value O) int

	Get(ix int) (O, error)

	Insert(value O, ix int) error
	Replace(value O, ix int) (O, error)
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

func Sort[O objects.ComparableObject[O]](list List[O]) {
	SortT(list, objects.ComparableComparator[O]())
}

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

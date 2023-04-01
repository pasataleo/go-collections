package collections

import (
	"testing"

	"github.com/pasataleo/go-objects/objects"
)

func TestArrayList_Collection(t *testing.T) {
	runCollectionTests(t, func() Collection[objects.String] {
		return NewArrayList[objects.String]()
	}, map[string]objects.String{
		"one":   objects.WrapString("one"),
		"two":   objects.WrapString("two"),
		"three": objects.WrapString("three"),
	})
}

func TestArrayList_List(t *testing.T) {
	runListTests(t, func() List[objects.String] {
		return NewArrayList[objects.String]()
	}, map[string]objects.String{
		"zero":  objects.WrapString("zero"),
		"one":   objects.WrapString("one"),
		"two":   objects.WrapString("two"),
		"three": objects.WrapString("three"),
		"four":  objects.WrapString("four"),
	})
}

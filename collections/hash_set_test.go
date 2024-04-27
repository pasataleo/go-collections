package collections

import (
	"testing"

	"github.com/pasataleo/go-objects/objects"
)

func TestHashSet_Collection(t *testing.T) {
	runCollectionTests(t, func() Collection[*objects.String] {
		return NewHashSet[*objects.String]()
	}, map[string]*objects.String{
		"one":   objects.WrapString("one"),
		"two":   objects.WrapString("two"),
		"three": objects.WrapString("three"),
	})
}

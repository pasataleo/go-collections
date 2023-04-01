package collections

import (
	"testing"

	"github.com/pasataleo/go-objects/objects"
)

func TestHashMap_Collection(t *testing.T) {
	makeEntry := func(key, value string) MapEntry[objects.String, objects.String] {
		return mapEntry[objects.String, objects.String]{
			key:            objects.WrapString(key),
			value:          objects.WrapString(value),
			keyConverter:   objects.ObjectIdentityConverter[objects.String](),
			valueConverter: objects.ObjectIdentityConverter[objects.String](),
		}
	}

	runCollectionTests(t, func() Collection[MapEntry[objects.String, objects.String]] {
		return NewHashMap[objects.String, objects.String]()
	}, map[string]MapEntry[objects.String, objects.String]{
		"one":   makeEntry("one", "four"),
		"two":   makeEntry("two", "five"),
		"three": makeEntry("three", "six"),
	})
}

func TestHashMap_Map(t *testing.T) {
	runMapTests(t, func() Map[objects.String, objects.String] {
		return NewHashMap[objects.String, objects.String]()
	}, map[string]objects.String{
		"zero":  objects.WrapString("zero"),
		"one":   objects.WrapString("one"),
		"two":   objects.WrapString("two"),
		"three": objects.WrapString("three"),
		"four":  objects.WrapString("four"),
		"five":  objects.WrapString("five"),
	})
}

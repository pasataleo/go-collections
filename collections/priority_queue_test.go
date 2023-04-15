package collections

import (
	"testing"

	"github.com/pasataleo/go-objects/objects"
	"github.com/pasataleo/go-testing/tests"
)

func TestHeap_Collection(t *testing.T) {
	runCollectionTests(t, func() Collection[objects.String] {
		return NewPriorityQueue[objects.String]()
	}, map[string]objects.String{
		"one":   objects.WrapString("one"),
		"two":   objects.WrapString("two"),
		"three": objects.WrapString("three"),
	})
}

func TestHeap(t *testing.T) {
	heap := NewPriorityQueue[objects.String]()

	tests.ExecFn(t, heap.Offer, objects.WrapString("a")).NoError()
	tests.ExecFn(t, heap.Offer, objects.WrapString("c")).NoError()
	tests.ExecFn(t, heap.Offer, objects.WrapString("b")).NoError()
	tests.ExecFn(t, heap.Size).Equals(3)

	tests.ExecFn(t, heap.Peep).NoError().ObjectEquals(objects.WrapString("a"))
	tests.ExecFn(t, heap.Pop).NoError().ObjectEquals(objects.WrapString("a"))
	tests.ExecFn(t, heap.Size).Equals(2)

	tests.ExecFn(t, heap.Peep).NoError().ObjectEquals(objects.WrapString("b"))
	tests.ExecFn(t, heap.Pop).NoError().ObjectEquals(objects.WrapString("b"))
	tests.ExecFn(t, heap.Size).Equals(1)

	tests.ExecFn(t, heap.Peep).NoError().ObjectEquals(objects.WrapString("c"))
	tests.ExecFn(t, heap.Pop).NoError().ObjectEquals(objects.WrapString("c"))
	tests.ExecFn(t, heap.Size).Equals(0)
}

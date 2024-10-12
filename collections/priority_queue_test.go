package collections

import (
	"testing"

	"github.com/pasataleo/go-objects/objects"
	"github.com/pasataleo/go-testing/tests"
)

func TestHeap_Collection(t *testing.T) {
	runCollectionTests(t, func() Collection[*objects.String] {
		return NewPriorityQueue[*objects.String]()
	}, map[string]*objects.String{
		"one":   objects.WrapString("one"),
		"two":   objects.WrapString("two"),
		"three": objects.WrapString("three"),
	})
}

func TestHeap(t *testing.T) {
	heap := NewPriorityQueue[*objects.String]()

	tests.ExecuteE(heap.Offer(objects.WrapString("a"))).NoError(t)
	tests.ExecuteE(heap.Offer(objects.WrapString("c"))).NoError(t)
	tests.ExecuteE(heap.Offer(objects.WrapString("b"))).NoError(t)
	tests.Execute(heap.Size()).Equal(t, 3)

	tests.Execute2E(heap.Peep()).NoError(t).Equal(t, objects.WrapString("a"))
	tests.Execute2E(heap.Pop()).NoError(t).Equal(t, objects.WrapString("a"))
	tests.Execute(heap.Size()).Equal(t, 2)

	tests.Execute2E(heap.Peep()).NoError(t).Equal(t, objects.WrapString("b"))
	tests.Execute2E(heap.Pop()).NoError(t).Equal(t, objects.WrapString("b"))
	tests.Execute(heap.Size()).Equal(t, 1)

	tests.Execute2E(heap.Peep()).NoError(t).Equal(t, objects.WrapString("c"))
	tests.Execute2E(heap.Pop()).NoError(t).Equal(t, objects.WrapString("c"))
	tests.Execute(heap.Size()).Equal(t, 0)
}

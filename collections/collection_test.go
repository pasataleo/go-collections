package collections

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pasataleo/go-errors/errors"
	"github.com/pasataleo/go-objects/objects"
	"github.com/pasataleo/go-objects/objects/cmpobjects"
	"github.com/pasataleo/go-testing/tests"
)

func runCollectionTests[V objects.Object](t *testing.T, init func() Collection[V], data map[string]V) {
	t.Run("collection_single", func(t *testing.T) {
		collection := init()

		tests.ExecuteE(collection.Add(data["one"])).NoError(t)
		tests.Execute(collection.Contains(data["one"])).Equal(t, true)
		tests.Execute(collection.Size()).Equal(t, 1)
		tests.ExecuteE(collection.Remove(data["one"])).NoError(t)
	})

	t.Run("collection_multi", func(t *testing.T) {
		starter := init()

		tests.ExecuteE(starter.Add(data["one"])).NoError(t)
		tests.ExecuteE(starter.Add(data["two"])).NoError(t)
		tests.ExecuteE(starter.Add(data["three"])).NoError(t)

		collection := init()
		tests.ExecuteE(collection.AddAll(starter)).NoError(t)
		tests.Execute(collection.ContainsAll(starter)).Equal(t, true)
		tests.ExecuteE(collection.RemoveAll(starter)).NoError(t)
		tests.Execute(collection.ContainsAll(starter)).Equal(t, false)

		err := tests.ExecuteE(collection.RemoveAll(starter)).Capture()
		tests.Execute(errors.GetErrorCode(err)).Equal(t, errors.ErrorCodeMulti, tests.Fatal)

		errs := errors.Expand(err)
		tests.Execute(len(errs)).Equal(t, 3)
	})

	t.Run("collection_failed", func(t *testing.T) {
		collection := init()

		tests.ExecuteE(collection.Add(data["one"])).NoError(t)
		tests.Execute(collection.Contains(data["two"])).Equal(t, false)
		tests.ExecuteE(collection.Remove(data["two"])).ErrorCode(t, ErrorCodeNotFound)
	})

	t.Run("cmp", func(t *testing.T) {
		one := init()
		tests.ExecuteE(one.Add(data["one"])).NoError(t)
		tests.ExecuteE(one.Add(data["two"])).NoError(t)
		tests.ExecuteE(one.Add(data["three"])).NoError(t)

		two := init()
		tests.ExecuteE(two.Add(data["one"])).NoError(t)
		tests.ExecuteE(two.Add(data["three"])).NoError(t)

		diff := cmp.Diff(one, two, cmpobjects.ObjectEquals())
		if len(diff) == 0 {
			t.Errorf("expected diff, got empty")
		}
		t.Logf("diff: %v", diff)
	})
}

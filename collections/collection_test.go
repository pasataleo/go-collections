package collections

import (
	"testing"

	"github.com/pasataleo/go-errors/errors"
	"github.com/pasataleo/go-objects/objects"
	"github.com/pasataleo/go-testing/tests"
)

func runCollectionTests[V objects.Object](t *testing.T, init func() Collection[V], data map[string]V) {
	t.Run("collection_single", func(t *testing.T) {
		collection := init()

		tests.ExecFn(t, collection.Add, data["one"]).NoError()
		tests.ExecFn(t, collection.Contains, data["one"]).True()
		tests.ExecFn(t, collection.Size).Equals(1)
		tests.ExecFn(t, collection.Remove, data["one"]).NoError()
	})

	t.Run("collection_multi", func(t *testing.T) {
		starter := init()

		tests.ExecFn(t, starter.Add, data["one"]).NoError()
		tests.ExecFn(t, starter.Add, data["two"]).NoError()
		tests.ExecFn(t, starter.Add, data["three"]).NoError()

		collection := init()
		tests.ExecFn(t, collection.AddAll, starter).NoError()
		tests.ExecFn(t, collection.ContainsAll, starter).True()
		tests.ExecFn(t, collection.RemoveAll, starter).NoError()
		tests.ExecFn(t, collection.ContainsAll, starter).False()

		capture, _ := tests.ExecFn(t, collection.RemoveAll, starter).CaptureError()
		capture.Fatal().ErrorCode(errors.ErrorCodeMulti)

		multi, _ := tests.ExecFnT[[]error](t, errors.GetEmbeddedData[[]error], capture.Value).
			True().
			Capture()
		multi.Len(3)
	})

	t.Run("collection_failed", func(t *testing.T) {
		collection := init()

		tests.ExecFn(t, collection.Add, data["one"]).NoError()
		tests.ExecFn(t, collection.Contains, data["two"]).False()
		tests.ExecFn(t, collection.Remove, data["two"]).ErrorCode(ErrorCodeNotFound)
	})
}

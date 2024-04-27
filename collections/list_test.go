package collections

import (
	"testing"

	"github.com/pasataleo/go-objects/objects"
	"github.com/pasataleo/go-testing/tests"
)

func runListTests(t *testing.T, init func() List[*objects.String], data map[string]*objects.String) {
	t.Run("list_insert", func(t *testing.T) {
		list := init()

		tests.ExecFn(t, list.Insert, data["zero"], 0).NoError()
		tests.ExecFn(t, list.Insert, data["one"], 1).NoError()
		tests.ExecFn(t, list.Insert, data["two"], 2).NoError()

		tests.ExecFn(t, list.Size).Equals(3)
		tests.ExecFn(t, list.Get, 0).NoError().ObjectEquals(data["zero"])
		tests.ExecFn(t, list.Get, 1).NoError().ObjectEquals(data["one"])
		tests.ExecFn(t, list.Get, 2).NoError().ObjectEquals(data["two"])

		tests.ExecFn(t, list.Insert, data["three"], 4).ErrorCode(ErrorCodeOutOfBounds)
		tests.ExecFn(t, list.Insert, data["three"], -1).ErrorCode(ErrorCodeOutOfBounds)

		tests.ExecFn(t, list.Insert, data["three"], 1)
		tests.ExecFn(t, list.Insert, data["four"], 2)

		tests.ExecFn(t, list.Size).Equals(5)
		tests.ExecFn(t, list.Get, 0).NoError().ObjectEquals(data["zero"])
		tests.ExecFn(t, list.Get, 1).NoError().ObjectEquals(data["three"])
		tests.ExecFn(t, list.Get, 2).NoError().ObjectEquals(data["four"])
		tests.ExecFn(t, list.Get, 3).NoError().ObjectEquals(data["one"])
		tests.ExecFn(t, list.Get, 4).NoError().ObjectEquals(data["two"])
	})

	t.Run("list_replace", func(t *testing.T) {
		list := init()

		tests.ExecFn(t, list.Insert, data["zero"], 0).NoError()
		tests.ExecFn(t, list.Insert, data["one"], 1).NoError()
		tests.ExecFn(t, list.Insert, data["two"], 2).NoError()

		tests.ExecFn(t, list.Size).Equals(3)
		tests.ExecFn(t, list.Get, 0).NoError().ObjectEquals(data["zero"])
		tests.ExecFn(t, list.Get, 1).NoError().ObjectEquals(data["one"])
		tests.ExecFn(t, list.Get, 2).NoError().ObjectEquals(data["two"])

		tests.ExecFn(t, list.Replace, data["two"], 0).NoError()
		tests.ExecFn(t, list.Replace, data["three"], 2).NoError()

		tests.ExecFn(t, list.Size).Equals(3)
		tests.ExecFn(t, list.Get, 0).NoError().ObjectEquals(data["two"])
		tests.ExecFn(t, list.Get, 1).NoError().ObjectEquals(data["one"])
		tests.ExecFn(t, list.Get, 2).NoError().ObjectEquals(data["three"])

		tests.ExecFn(t, list.Replace, data["four"], -1).ErrorCode(ErrorCodeOutOfBounds)
		tests.ExecFn(t, list.Replace, data["four"], 4).ErrorCode(ErrorCodeOutOfBounds)
	})

	t.Run("list_remove", func(t *testing.T) {
		list := init()

		tests.ExecFn(t, list.Insert, data["zero"], 0).NoError()
		tests.ExecFn(t, list.Insert, data["one"], 1).NoError()
		tests.ExecFn(t, list.Insert, data["two"], 2).NoError()

		tests.ExecFn(t, list.Size).Equals(3)
		tests.ExecFn(t, list.RemoveAt, 1).NoError()
		tests.ExecFn(t, list.Get, 1).NoError().ObjectEquals(data["two"])
	})

	t.Run("list_index", func(t *testing.T) {
		list := init()

		tests.ExecFn(t, list.Insert, data["zero"], 0).NoError()
		tests.ExecFn(t, list.Insert, data["one"], 1).NoError()
		tests.ExecFn(t, list.Insert, data["two"], 2).NoError()

		tests.ExecFn(t, list.IndexOf, data["zero"]).Equals(0)
		tests.ExecFn(t, list.IndexOf, data["one"]).Equals(1)
		tests.ExecFn(t, list.IndexOf, data["two"]).Equals(2)
	})
}

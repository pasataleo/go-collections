package collections

import (
	"testing"

	"github.com/pasataleo/go-objects/objects"
	"github.com/pasataleo/go-testing/tests"
)

func runListTests(t *testing.T, init func() List[*objects.String], data map[string]*objects.String) {
	t.Run("list_insert", func(t *testing.T) {
		list := init()

		tests.ExecuteE(list.Insert(data["zero"], 0)).NoError(t)
		tests.ExecuteE(list.Insert(data["one"], 1)).NoError(t)
		tests.ExecuteE(list.Insert(data["two"], 2)).NoError(t)

		tests.Execute(list.Size()).Equal(t, 3)
		tests.Execute2E(list.Get(0)).NoError(t).Equal(t, data["zero"])
		tests.Execute2E(list.Get(1)).NoError(t).Equal(t, data["one"])
		tests.Execute2E(list.Get(2)).NoError(t).Equal(t, data["two"])

		tests.ExecuteE(list.Insert(data["three"], 4)).ErrorCode(t, ErrorCodeOutOfBounds)
		tests.ExecuteE(list.Insert(data["three"], -1)).ErrorCode(t, ErrorCodeOutOfBounds)

		tests.ExecuteE(list.Insert(data["three"], 1)).NoError(t)
		tests.ExecuteE(list.Insert(data["four"], 2)).NoError(t)

		tests.Execute(list.Size()).Equal(t, 5)
		tests.Execute2E(list.Get(0)).NoError(t).Equal(t, data["zero"])
		tests.Execute2E(list.Get(1)).NoError(t).Equal(t, data["three"])
		tests.Execute2E(list.Get(2)).NoError(t).Equal(t, data["four"])
		tests.Execute2E(list.Get(3)).NoError(t).Equal(t, data["one"])
		tests.Execute2E(list.Get(4)).NoError(t).Equal(t, data["two"])
	})

	t.Run("list_replace", func(t *testing.T) {
		list := init()

		tests.ExecuteE(list.Insert(data["zero"], 0)).NoError(t)
		tests.ExecuteE(list.Insert(data["one"], 1)).NoError(t)
		tests.ExecuteE(list.Insert(data["two"], 2)).NoError(t)

		tests.Execute(list.Size()).Equal(t, 3)
		tests.Execute2E(list.Get(0)).NoError(t).Equal(t, data["zero"])
		tests.Execute2E(list.Get(1)).NoError(t).Equal(t, data["one"])
		tests.Execute2E(list.Get(2)).NoError(t).Equal(t, data["two"])

		tests.Execute2E(list.Replace(data["two"], 0)).NoError(t)
		tests.Execute2E(list.Replace(data["one"], 1)).NoError(t)

		tests.Execute(list.Size()).Equal(t, 3)
		tests.Execute2E(list.Get(0)).NoError(t).Equal(t, data["two"])
		tests.Execute2E(list.Get(1)).NoError(t).Equal(t, data["one"])
		tests.Execute2E(list.Get(2)).NoError(t).Equal(t, data["two"])

		tests.Execute2E(list.Replace(data["four"], -1)).ErrorCode(t, ErrorCodeOutOfBounds)
		tests.Execute2E(list.Replace(data["four"], 4)).ErrorCode(t, ErrorCodeOutOfBounds)
	})

	t.Run("list_remove", func(t *testing.T) {
		list := init()

		tests.ExecuteE(list.Insert(data["zero"], 0)).NoError(t)
		tests.ExecuteE(list.Insert(data["one"], 1)).NoError(t)
		tests.ExecuteE(list.Insert(data["two"], 2)).NoError(t)

		tests.Execute(list.Size()).Equal(t, 3)
		tests.Execute2E(list.RemoveAt(1)).NoError(t).Equal(t, data["one"])
		tests.Execute2E(list.Get(1)).NoError(t).Equal(t, data["two"])
	})

	t.Run("list_index", func(t *testing.T) {
		list := init()

		tests.ExecuteE(list.Insert(data["zero"], 0)).NoError(t)
		tests.ExecuteE(list.Insert(data["one"], 1)).NoError(t)
		tests.ExecuteE(list.Insert(data["two"], 2)).NoError(t)

		tests.Execute(list.IndexOf(data["zero"])).Equal(t, 0)
		tests.Execute(list.IndexOf(data["one"])).Equal(t, 1)
		tests.Execute(list.IndexOf(data["two"])).Equal(t, 2)
	})
}

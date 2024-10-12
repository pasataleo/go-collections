package collections

import (
	"testing"

	"github.com/pasataleo/go-objects/objects"
	"github.com/pasataleo/go-testing/tests"
)

func runMapTests(t *testing.T, init func() Map[*objects.String, *objects.String], data map[string]*objects.String) {
	t.Run("map_put", func(t *testing.T) {
		collection := init()

		tests.ExecuteE(collection.Put(data["zero"], data["three"])).NoError(t)
		tests.ExecuteE(collection.Put(data["one"], data["four"])).NoError(t)
		tests.ExecuteE(collection.Put(data["two"], data["five"])).NoError(t)
		tests.ExecuteE(collection.Put(data["two"], data["four"])).ErrorCode(t, ErrorCodeAlreadyExists)

		tests.Execute(collection.ContainsKey(data["zero"])).Equal(t, true)
		tests.Execute(collection.ContainsKey(data["one"])).Equal(t, true)
		tests.Execute(collection.ContainsKey(data["two"])).Equal(t, true)

		tests.Execute(collection.Get(data["zero"])).Equal(t, data["three"])
		tests.Execute(collection.Get(data["one"])).Equal(t, data["four"])
		tests.Execute(collection.Get(data["two"])).Equal(t, data["five"])
	})
	t.Run("map_replace", func(t *testing.T) {
		collection := init()

		tests.ExecuteE(collection.Put(data["zero"], data["three"])).NoError(t)
		tests.ExecuteE(collection.Put(data["one"], data["four"])).NoError(t)
		tests.ExecuteE(collection.Put(data["two"], data["five"])).NoError(t)

		tests.Execute2E(collection.Replace(data["zero"], data["four"])).NoError(t).Equal(t, data["three"])
		tests.Execute2E(collection.Replace(data["two"], data["four"])).NoError(t).Equal(t, data["five"])
		tests.Execute2E(collection.Replace(data["four"], data["four"])).ErrorCode(t, ErrorCodeNotFound)

		tests.Execute(collection.Get(data["zero"])).Equal(t, data["four"])
		tests.Execute(collection.Get(data["one"])).Equal(t, data["four"])
		tests.Execute(collection.Get(data["two"])).Equal(t, data["four"])
	})
	t.Run("map_put_or_replace", func(t *testing.T) {
		collection := init()

		tests.ExecuteE(collection.Put(data["zero"], data["three"])).NoError(t)
		tests.ExecuteE(collection.Put(data["one"], data["four"])).NoError(t)
		tests.ExecuteE(collection.Put(data["two"], data["five"])).NoError(t)

		tests.Execute2(collection.PutOrReplace(data["zero"], data["four"])).Equal(t, true).Equal(t, data["three"])
		tests.Execute2(collection.PutOrReplace(data["two"], data["four"])).Equal(t, true).Equal(t, data["five"])
		tests.Execute2(collection.PutOrReplace(data["four"], data["four"])).Equal(t, false)

		tests.Execute(collection.Get(data["zero"])).Equal(t, data["four"])
		tests.Execute(collection.Get(data["one"])).Equal(t, data["four"])
		tests.Execute(collection.Get(data["two"])).Equal(t, data["four"])
		tests.Execute(collection.Get(data["four"])).Equal(t, data["four"])
	})
	t.Run("map_delete", func(t *testing.T) {
		collection := init()

		tests.ExecuteE(collection.Put(data["zero"], data["three"])).NoError(t)
		tests.ExecuteE(collection.Put(data["one"], data["four"])).NoError(t)
		tests.ExecuteE(collection.Put(data["two"], data["five"])).NoError(t)

		tests.Execute(collection.ContainsKey(data["zero"])).Equal(t, true)
		tests.Execute(collection.ContainsKey(data["one"])).Equal(t, true)
		tests.Execute(collection.ContainsKey(data["two"])).Equal(t, true)

		tests.Execute2E(collection.Delete(data["zero"])).NoError(t).Equal(t, data["three"])
		tests.Execute2E(collection.Delete(data["one"])).NoError(t).Equal(t, data["four"])
		tests.Execute2E(collection.Delete(data["two"])).NoError(t).Equal(t, data["five"])

		tests.Execute(collection.ContainsKey(data["zero"])).Equal(t, false)
		tests.Execute(collection.ContainsKey(data["one"])).Equal(t, false)
		tests.Execute(collection.ContainsKey(data["two"])).Equal(t, false)

		tests.Execute2E(collection.Delete(data["zero"])).ErrorCode(t, ErrorCodeNotFound)
	})
	t.Run("map_delete_if_present", func(t *testing.T) {
		collection := init()

		tests.ExecuteE(collection.Put(data["zero"], data["three"])).NoError(t)
		tests.ExecuteE(collection.Put(data["one"], data["four"])).NoError(t)
		tests.ExecuteE(collection.Put(data["two"], data["five"])).NoError(t)

		tests.Execute(collection.ContainsKey(data["zero"])).Equal(t, true)
		tests.Execute(collection.ContainsKey(data["one"])).Equal(t, true)
		tests.Execute(collection.ContainsKey(data["two"])).Equal(t, true)

		tests.Execute2(collection.DeleteIfPresent(data["zero"])).Equal(t, true).Equal(t, data["three"])
		tests.Execute2(collection.DeleteIfPresent(data["one"])).Equal(t, true).Equal(t, data["four"])
		tests.Execute2(collection.DeleteIfPresent(data["two"])).Equal(t, true).Equal(t, data["five"])

		tests.Execute(collection.ContainsKey(data["zero"])).Equal(t, false)
		tests.Execute(collection.ContainsKey(data["one"])).Equal(t, false)
		tests.Execute(collection.ContainsKey(data["two"])).Equal(t, false)

		tests.Execute2(collection.DeleteIfPresent(data["zero"])).Equal(t, false)
	})
	t.Run("map_get", func(t *testing.T) {
		collection := init()

		tests.ExecuteE(collection.Put(data["zero"], data["three"])).NoError(t)
		tests.ExecuteE(collection.Put(data["one"], data["four"])).NoError(t)
		tests.ExecuteE(collection.Put(data["two"], data["five"])).NoError(t)

		tests.Execute2E(collection.GetSafe(data["zero"])).NoError(t).Equal(t, data["three"])
		tests.Execute2E(collection.GetSafe(data["one"])).NoError(t).Equal(t, data["four"])
		tests.Execute2E(collection.GetSafe(data["two"])).NoError(t).Equal(t, data["five"])
		tests.Execute2E(collection.GetSafe(data["three"])).ErrorCode(t, ErrorCodeNotFound)
	})
	t.Run("map_keys", func(t *testing.T) {
		collection := init()

		tests.ExecuteE(collection.Put(data["zero"], data["three"])).NoError(t)
		tests.ExecuteE(collection.Put(data["one"], data["four"])).NoError(t)
		tests.ExecuteE(collection.Put(data["two"], data["five"])).NoError(t)

		keys := collection.Keys()
		tests.Execute(keys.Size()).Equal(t, 3)
		tests.Execute(keys.Contains(data["zero"])).Equal(t, true)
		tests.Execute(keys.Contains(data["one"])).Equal(t, true)
		tests.Execute(keys.Contains(data["two"])).Equal(t, true)
	})
	t.Run("map_values", func(t *testing.T) {
		collection := init()

		tests.ExecuteE(collection.Put(data["zero"], data["three"])).NoError(t)
		tests.ExecuteE(collection.Put(data["one"], data["four"])).NoError(t)
		tests.ExecuteE(collection.Put(data["two"], data["five"])).NoError(t)

		values := collection.Values()
		tests.Execute(values.Size()).Equal(t, 3)
		tests.Execute(values.Contains(data["three"])).Equal(t, true)
		tests.Execute(values.Contains(data["four"])).Equal(t, true)
		tests.Execute(values.Contains(data["five"])).Equal(t, true)
	})
}

package collections

import (
	"testing"

	"github.com/pasataleo/go-objects/objects"
	"github.com/pasataleo/go-testing/tests"
)

func runMapTests(t *testing.T, init func() Map[objects.String, objects.String], data map[string]objects.String) {
	t.Run("map_put", func(t *testing.T) {
		collection := init()

		tests.ExecFn(t, collection.Put, data["zero"], data["three"]).NoError()
		tests.ExecFn(t, collection.Put, data["one"], data["four"]).NoError()
		tests.ExecFn(t, collection.Put, data["two"], data["five"]).NoError()
		tests.ExecFn(t, collection.Put, data["two"], data["four"]).ErrorCode(ErrorCodeAlreadyExists)

		tests.ExecFn(t, collection.ContainsKey, data["zero"]).True()
		tests.ExecFn(t, collection.ContainsKey, data["one"]).True()
		tests.ExecFn(t, collection.ContainsKey, data["two"]).True()

		tests.ExecFn(t, collection.Get, data["zero"]).ObjectEquals(data["three"])
		tests.ExecFn(t, collection.Get, data["one"]).ObjectEquals(data["four"])
		tests.ExecFn(t, collection.Get, data["two"]).ObjectEquals(data["five"])
	})
	t.Run("map_replace", func(t *testing.T) {
		collection := init()

		tests.ExecFn(t, collection.Put, data["zero"], data["three"]).NoError()
		tests.ExecFn(t, collection.Put, data["one"], data["four"]).NoError()
		tests.ExecFn(t, collection.Put, data["two"], data["five"]).NoError()

		tests.ExecFn(t, collection.Replace, data["zero"], data["four"]).NoError()
		tests.ExecFn(t, collection.Replace, data["two"], data["four"]).NoError()
		tests.ExecFn(t, collection.Replace, data["four"], data["four"]).ErrorCode(ErrorCodeNotFound)

		tests.ExecFn(t, collection.Get, data["zero"]).ObjectEquals(data["four"])
		tests.ExecFn(t, collection.Get, data["one"]).ObjectEquals(data["four"])
		tests.ExecFn(t, collection.Get, data["two"]).ObjectEquals(data["four"])
	})
	t.Run("map_put_or_replace", func(t *testing.T) {
		collection := init()

		tests.ExecFn(t, collection.Put, data["zero"], data["three"]).NoError()
		tests.ExecFn(t, collection.Put, data["one"], data["four"]).NoError()
		tests.ExecFn(t, collection.Put, data["two"], data["five"]).NoError()

		tests.ExecFn(t, collection.PutOrReplace, data["zero"], data["four"]).True().ObjectEquals(data["three"])
		tests.ExecFn(t, collection.PutOrReplace, data["two"], data["four"]).True().ObjectEquals(data["five"])
		tests.ExecFn(t, collection.PutOrReplace, data["four"], data["four"]).False()

		tests.ExecFn(t, collection.Get, data["zero"]).ObjectEquals(data["four"])
		tests.ExecFn(t, collection.Get, data["one"]).ObjectEquals(data["four"])
		tests.ExecFn(t, collection.Get, data["two"]).ObjectEquals(data["four"])
		tests.ExecFn(t, collection.Get, data["four"]).ObjectEquals(data["four"])
	})
	t.Run("map_delete", func(t *testing.T) {
		collection := init()

		tests.ExecFn(t, collection.Put, data["zero"], data["three"]).NoError()
		tests.ExecFn(t, collection.Put, data["one"], data["four"]).NoError()
		tests.ExecFn(t, collection.Put, data["two"], data["five"]).NoError()

		tests.ExecFn(t, collection.ContainsKey, data["zero"]).True()
		tests.ExecFn(t, collection.ContainsKey, data["one"]).True()
		tests.ExecFn(t, collection.ContainsKey, data["two"]).True()

		tests.ExecFn(t, collection.Delete, data["zero"]).NoError().ObjectEquals(data["three"])
		tests.ExecFn(t, collection.Delete, data["one"]).NoError().ObjectEquals(data["four"])
		tests.ExecFn(t, collection.Delete, data["two"]).NoError().ObjectEquals(data["five"])

		tests.ExecFn(t, collection.ContainsKey, data["zero"]).False()
		tests.ExecFn(t, collection.ContainsKey, data["one"]).False()
		tests.ExecFn(t, collection.ContainsKey, data["two"]).False()

		tests.ExecFn(t, collection.Delete, data["zero"]).ErrorCode(ErrorCodeNotFound)
	})
	t.Run("map_delete_if_present", func(t *testing.T) {
		collection := init()

		tests.ExecFn(t, collection.Put, data["zero"], data["three"]).NoError()
		tests.ExecFn(t, collection.Put, data["one"], data["four"]).NoError()
		tests.ExecFn(t, collection.Put, data["two"], data["five"]).NoError()

		tests.ExecFn(t, collection.ContainsKey, data["zero"]).True()
		tests.ExecFn(t, collection.ContainsKey, data["one"]).True()
		tests.ExecFn(t, collection.ContainsKey, data["two"]).True()

		tests.ExecFn(t, collection.DeleteIfPresent, data["zero"]).True().ObjectEquals(data["three"])
		tests.ExecFn(t, collection.DeleteIfPresent, data["one"]).True().ObjectEquals(data["four"])
		tests.ExecFn(t, collection.DeleteIfPresent, data["two"]).True().ObjectEquals(data["five"])

		tests.ExecFn(t, collection.ContainsKey, data["zero"]).False()
		tests.ExecFn(t, collection.ContainsKey, data["one"]).False()
		tests.ExecFn(t, collection.ContainsKey, data["two"]).False()

		tests.ExecFn(t, collection.DeleteIfPresent, data["zero"]).False()
	})
	t.Run("map_get", func(t *testing.T) {
		collection := init()

		tests.ExecFn(t, collection.Put, data["zero"], data["three"]).NoError()
		tests.ExecFn(t, collection.Put, data["one"], data["four"]).NoError()
		tests.ExecFn(t, collection.Put, data["two"], data["five"]).NoError()

		tests.ExecFn(t, collection.GetSafe, data["zero"]).NoError().ObjectEquals(data["three"])
		tests.ExecFn(t, collection.GetSafe, data["one"]).NoError().ObjectEquals(data["four"])
		tests.ExecFn(t, collection.GetSafe, data["two"]).NoError().ObjectEquals(data["five"])
		tests.ExecFn(t, collection.GetSafe, data["three"]).ErrorCode(ErrorCodeNotFound)
	})
	t.Run("map_keys", func(t *testing.T) {
		collection := init()

		tests.ExecFn(t, collection.Put, data["zero"], data["three"]).NoError()
		tests.ExecFn(t, collection.Put, data["one"], data["four"]).NoError()
		tests.ExecFn(t, collection.Put, data["two"], data["five"]).NoError()

		keys := collection.Keys()
		tests.ExecFn(t, keys.Size).Equals(3)
		tests.ExecFn(t, keys.Contains, data["zero"]).True()
		tests.ExecFn(t, keys.Contains, data["one"]).True()
		tests.ExecFn(t, keys.Contains, data["two"]).True()
	})
	t.Run("map_values", func(t *testing.T) {
		collection := init()

		tests.ExecFn(t, collection.Put, data["zero"], data["three"]).NoError()
		tests.ExecFn(t, collection.Put, data["one"], data["four"]).NoError()
		tests.ExecFn(t, collection.Put, data["two"], data["five"]).NoError()

		values := collection.Values()
		tests.ExecFn(t, values.Size).Equals(3)
		tests.ExecFn(t, values.Contains, data["three"]).True()
		tests.ExecFn(t, values.Contains, data["four"]).True()
		tests.ExecFn(t, values.Contains, data["five"]).True()
	})
}

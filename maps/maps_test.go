package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("key exists", func(t *testing.T) {
		got, err := dict.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
		assertError(t, err, nil)
	})
	t.Run("key doesn't exist", func(t *testing.T) {
		_, got := dict.Search("mykey")
		assertError(t, got, ErrMissingKey)
	})

}

func TestAdd(t *testing.T) {
	dict := Dictionary{}

	t.Run("add new key", func(t *testing.T) {
		errExists := dict.Add("key1", "value1")
		got, err := dict.Search("key1")
		want := "value1"
		assertString(t, got, want)
		assertError(t, err, nil)
		assertError(t, errExists, nil)
	})
	t.Run("add existing key", func(t *testing.T) {
		err := dict.Add("key1", "value1")
		assertError(t, err, ErrKeyAlreadyExists)
	})
}

func TestUpdate(t *testing.T) {
	dict := Dictionary{"key1": "value1"}

	t.Run("update existing key", func(t *testing.T) {
		err := dict.Update("key1", "value1-updated")
		got, _ := dict.Search("key1")
		expected := "value1-updated"
		assertString(t, got, expected)
		assertError(t, err, nil)

	})
	t.Run("try updating missing key", func(t *testing.T) {
		errExists := dict.Update("key2", "value2")
		assertError(t, errExists, ErrMissingKey)
	})
}

func TestDelete(t *testing.T) {
	dict := Dictionary{"key1": "value1"}

	t.Run("delete valid key", func(t *testing.T) {
		err := dict.Delete("key1")
		_, got := dict.Search("key1")
		assertError(t, err, nil)
		assertError(t, got, ErrMissingKey)
	})
	t.Run("delete invalid key", func(t *testing.T) {
		err := dict.Delete("key2")
		assertError(t, err, ErrMissingKey)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

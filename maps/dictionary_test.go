package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("word is in Dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)

	})

	t.Run("word is not in dictionary", func(t *testing.T) {
		dictionary := Dictionary{"hello": "hello stranger!"}
		_, got := dictionary.Search("test")

		assertError(t, got, ErrNotFound)
	})
}

func TestDictionary_Add(t *testing.T) {
	t.Run("adding a word to dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is a test"

		dictionary.Add(key, value)
		assertDefinition(t, dictionary, key, value)

	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, key string, value string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word")
	}
	assertStrings(t, got, value)

}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if !errors.Is(want, got) {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

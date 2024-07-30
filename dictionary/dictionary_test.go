package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("search", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is test"}
		want := "this is test"

		got := Search(dictionary, "test")

		if got != want {
			t.Errorf("want %q got %q", want, got)
		}
	})

	t.Run("dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is test"}
		want := "this is test"

		got := dictionary.Search("test")

		if got != want {
			t.Errorf("want %q got %q", want, got)
		}
	})

	t.Run("dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is test"}
		want := ""

		got := dictionary.Search("a,b,c,d")

		if got != want {
			t.Errorf("want %q got %q", want, got)
		}
	})
}

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

	t.Run("add", func(t *testing.T) {
		dictionary := Dictionary{}
		want := "this is test"

		dictionary.Add("test", "this is test")

		got := dictionary.Search("test")
		if got != want {
			t.Errorf("want %q got %q", want, got)
		}
	})

	t.Run("add vith error", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is test"}

		err := dictionary.Add("test", "this is new test")

		if err == nil {
			t.Errorf("want error but got nil")
		}
	})

	t.Run("update", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is test"}
		want := "this is new test"

		dictionary.Update("test", "this is new test")

		got := dictionary.Search("test")
		if got != want {
			t.Errorf("want %q got %q", want, got)
		}
	})

	t.Run("delete", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is test"}
		want := ""

		dictionary.Delete("test")

		got := dictionary.Search("test")
		if got != want {
			t.Errorf("want %q got %q", want, got)
		}
	})
}

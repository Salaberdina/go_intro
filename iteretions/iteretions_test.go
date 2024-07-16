package iteretions

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeat a", func(t *testing.T) {
		expected := "aaaaa"
		got := Repeat("a", 5)
		if expected != got {
			t.Errorf("expected %q got %q", expected, got)
		}
	})

	t.Run("Repeat b", func(t *testing.T) {
		expected := "bbbbb"
		got := Repeat("b", 5)
		if expected != got {
			t.Errorf("expected %q got %q", expected, got)
		}
	})

	t.Run("Repeat c 3 times", func(t *testing.T) {
		expected := "ccc"
		got := Repeat("c", 3)
		if expected != got {
			t.Errorf("expected %q got %q", expected, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1000000)
	}
}

package animals

import "testing"

func TestAnimals(t *testing.T) {
	t.Run("CatVoice", func(t *testing.T) {
		want := "may"
		c := Cat{
			Voice: "may",
		}
		got := c.Say()
		if want != got {
			t.Errorf("want %q got %q", want, got)
		}
	})
}

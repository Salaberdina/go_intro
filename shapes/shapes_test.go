package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Perimeter", func(t *testing.T) {
		r := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}
		want := 40.0
		got := Perimeter(r)
		if want != got {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("RectangleArea", func(t *testing.T) {
		r := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}
		want := 100.0
		got := r.Area()
		if want != got {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("CircleArea", func(t *testing.T) {
		c := Circle{
			Radius: 10.0,
		}
		want := 314.1592653589793
		got := c.Area()
		if want != got {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("RectangleArea", func(t *testing.T) {
		r := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}
		want := 100.0
		got := Area(r)
		if want != got {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("CircleArea", func(t *testing.T) {
		c := Circle{
			Radius: 10.0,
		}
		want := 314.1592653589793
		got := Area(c)
		if want != got {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

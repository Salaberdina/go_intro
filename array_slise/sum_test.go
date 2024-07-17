package arrayslise

import "testing"

func TestSum(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		want := 15
		got := Sum(numbers)
		if want != got {
			t.Errorf("want %d got %d, given %v", want, got, numbers)
		}
	})

	t.Run("array1", func(t *testing.T) {
		numbers := [5]int{5, 6, 7, 8, 9}
		want := 35
		got := Sum(numbers)
		if want != got {
			t.Errorf("want %d got %d, given %v", want, got, numbers)
		}
	})

	t.Run("slise", func(t *testing.T) {
		numbers := []int{5, 6, 7, 8, 9}
		want := 35
		got := SumSlise(numbers)
		if want != got {
			t.Errorf("want %d got %d, given %v", want, got, numbers)
		}
	})

	t.Run("slise", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		want := 6
		got := SumSlise(numbers)
		if want != got {
			t.Errorf("want %d got %d, given %v", want, got, numbers)
		}

	})
}

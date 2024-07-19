package arrayslise

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
	t.Run("Test SumAll", func(t *testing.T) {
		want := []int{3, 9}
		got := SumAll([]int{1, 2}, []int{0, 9})
		require.Equal(t, want, got)
	})

	t.Run("Test SumAll", func(t *testing.T) {
		want := []int{5, 4}
		got := SumAll([]int{3, 2}, []int{2, 2})
		require.Equal(t, want, got)
	})

	t.Run("Test SumAll", func(t *testing.T) {
		want := []int{5, 4, 6}
		got := SumAll([]int{3, 2}, []int{2, 2}, []int{3, 3})
		require.Equal(t, want, got)
	})

	t.Run("Test SumAll", func(t *testing.T) {
		want := []int{}
		got := SumAll()
		require.Equal(t, want, got)
	})

	t.Run("Test SumAll", func(t *testing.T) {
		want := []int{0}
		got := SumAll([]int{})
		require.Equal(t, want, got)
	})
}

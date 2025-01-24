package arraysslices

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	t.Run("array", func(t *testing.T) {
		got := Sum(arr)
		expected := 15
		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})

	slice := []int{1, 2, 3, 4, 5}
	t.Run("slices", func(t *testing.T) {
		got := Sum(slice)
		expected := 15
		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("non-empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, -1})
		want := []int{5, 8}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("one empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, -1})
		want := []int{0, 8}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

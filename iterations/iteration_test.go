package iterations

import (
	"testing"
)

func TestIterate(t *testing.T) {
	got := Iterate("a", 5)
	expected := "aaaaa"
	if got != expected {
		t.Errorf("got %q and want %q", got, expected)
	}
}

func BenchmarkIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a", 5)
	}
}

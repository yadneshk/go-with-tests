package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Yadnesh")
	got := buffer.String()
	expected := "Hello, Yadnesh"

	if got != expected {
		t.Errorf("got %q want %q expected", got, expected)
	}
}

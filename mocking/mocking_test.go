package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

type SpyCalledOperations struct {
	calls []string
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (spy *SpyCalledOperations) Sleep() {
	spy.calls = append(spy.calls, "sleep")
}

func (spy *SpyCalledOperations) Write(p []byte) (n int, err error) {
	spy.calls = append(spy.calls, "write")
	return
}

func TestCountdown(t *testing.T) {

	buffer := bytes.Buffer{}
	spySleeper := SpySleeper{}
	Countdown(&buffer, &spySleeper)

	got := buffer.String()
	expected := `3
2
1
Go!`
	if got != expected {
		t.Errorf("got %s want %s", got, expected)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough sleep calls, want 3 got %d", spySleeper.Calls)
	}

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCalledOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(want, spySleepPrinter.calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.calls)
		}
	})

}

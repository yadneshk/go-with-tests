package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

func main() {
	sleeper := DefaultSleeper{}
	Countdown(os.Stdout, &sleeper)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second * 3)
}

func Countdown(write io.Writer, sleeper Sleeper) {
	for i := 3; i >= 1; i-- {
		fmt.Fprintln(write, i)
		// time.Sleep(time.Second * 1)
		sleeper.Sleep()
	}
	fmt.Fprint(write, "Go!")
}

package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("print 3 up til Go!", func(t *testing.T) {
		var buffer bytes.Buffer
		spySleep := new(SpyCountdown)

		Countdown(&buffer, spySleep)

		got := buffer.String()
		expect := `3
2
1
Go!`

		if got != expect {
			t.Errorf("expected: %q, got: %q", expect, got)
		}
	})

	t.Run("sleep calls before every print", func(t *testing.T) {
		spySleepPrinter := new(SpyCountdown)
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := time.Second * 5

	spyTime := new(SpyTime)
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

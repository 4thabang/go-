package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type Sleeper interface {
	Sleep()
}

type SpyCountdown struct {
	Calls []string
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func main() {
	sleeper := &ConfigurableSleeper{
		duration: time.Second * 1,
		sleep:    time.Sleep,
	}
	Countdown(os.Stdout, sleeper)
}

func (s *SpyCountdown) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdown) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(w io.Writer, sleeper Sleeper) {
	values := []string{"3", "2", "1"}
	for _, value := range values {
		fmt.Fprintln(w, value)
		sleeper.Sleep()
	}
	fmt.Fprint(w, "Go!")
}

package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrTimeout = errors.New("10s timeout reached")
	timeout    = time.Second * 10
)

func main() {
	url := []string{
		"http://facebook.com",
		"http://techcrunch.com",
	}
	WebsiteRacer(url[0], url[1])
}

func WebsiteRacer(urlOne, urlTwo string) (string, error) {
	return ConfigurableRacer(urlOne, urlTwo, timeout)
}

func ConfigurableRacer(urlOne, urlTwo string, timeout time.Duration) (string, error) {
	for {
		select {
		case <-pingRacer(urlOne):
			return urlOne, nil
		case <-pingRacer(urlTwo):
			return urlTwo, nil
		case <-time.After(timeout):
			return "", fmt.Errorf("racer: %w", ErrTimeout)
		}
	}
}

func pingRacer(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

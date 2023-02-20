package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var timer = time.Millisecond * 20

func TestWebsiteRacer(t *testing.T) {
	t.Run("quickest server to respond", func(t *testing.T) {
		slowServer := WebsiteRacerHelper(time.Millisecond * 20)
		fastServer := WebsiteRacerHelper(time.Microsecond * 0)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer fastServer.Close()
		defer slowServer.Close()

		want := fastURL
		got, _ := WebsiteRacer(slowURL, fastURL)

		if got != want {
			t.Errorf("wanted: %s, got: %s", want, got)
		}
	})

	t.Run("return err if 10sec timeout is reached", func(t *testing.T) {
		server := WebsiteRacerHelper(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, timer)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func RacerHelper(t testing.TB, got, want string, err error) {
	t.Helper()
}

func WebsiteRacerHelper(delay time.Duration) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	})
	return httptest.NewServer(handler)
}

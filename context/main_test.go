package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, nil
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	chData := make(chan string, 1)

	go func() {
		var result string
		for _, res := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(time.Millisecond * 10)
				result += string(res)
			}
		}
		chData <- result
	}()

	select {
	case res := <-chData:
		return res, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "Hello, world"

	t.Run("cancel request with context", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		ctx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(time.Millisecond*5, cancel)
		req = req.WithContext(ctx)

		res := new(SpyResponseWriter)
		srv.ServeHTTP(res, req)

		if res.written {
			t.Error("a response should not have been written")
		}
	})

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		srv.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Errorf("got: %s, wanted: %s", res.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})
}

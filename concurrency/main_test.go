package main

import (
	"reflect"
	"testing"
	"time"
)

func mockCheckSites(url string) bool {
	return url != "https://witter.com"
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"https://google.com",
		"https://uber.com",
		"https://witter.com",
		"https://techcrunch.com",
	}

	want := map[string]bool{
		"https://google.com":     true,
		"https://uber.com":       true,
		"https://witter.com":     false,
		"https://techcrunch.com": true,
	}
	got := CheckWebsites(mockCheckSites, urls)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v, got: %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

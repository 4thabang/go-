package main

import "fmt"

type WebsiteChecker func(string) bool

type Result struct {
	url    string
	status bool
}

func main() {
	urls := []string{
		"https://google.com",
		"https://uber.com",
		"https://witter.com",
		"https://techcrunch.com",
	}
	checker := WebsiteChecker(CheckSites)
	siteStatus := CheckWebsites(checker, urls)

	for url, status := range siteStatus {
		fmt.Println(url, status)
	}
}

func CheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan Result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- Result{u, checker(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.url] = r.status
	}

	return results
}

func CheckSites(url string) bool {
	return url != "https://witter.com"
}

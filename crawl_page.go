package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
	}

	// skip other websites
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizedURL: %v\n", err)
	}

	isFirst := cfg.addPageVist(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Printf("Crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v\n", err)
	}

	nextURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v\n", err)
	}

	for _, nextURL := range nextURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}

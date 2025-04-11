package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch HTML: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("unexpected content type: %s", contentType)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read HTML body: %v", err)
	}

	htmlBody := string(body)

	return htmlBody, nil
}

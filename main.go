package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseURL := os.Args[1]

	fmt.Printf("starting crawl of: %v...\n", baseURL)
	htmlBody, err := getHTML(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(htmlBody)
}

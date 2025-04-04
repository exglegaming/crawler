package main

func main() {
	normalizeURL("https://blog.boot.dev/path/")
	normalizeURL("https://blog.boot.dev/path")
	normalizeURL("http://blog.boot.dev/path/")
	normalizeURL("https://blog.boot.dev/path")
}

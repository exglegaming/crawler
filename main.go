package main

const htmlBody = `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`

func main() {
	getURLsFromHTML(htmlBody, "https://blog.boot.dev")
}

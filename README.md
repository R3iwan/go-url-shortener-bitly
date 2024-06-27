# URL Shortener in Go

This is a simple command-line tool that takes a long URL as input, validates it, and generates a shortened URL using the Bitly API.

## Installation

1. Install Go (if not already installed) from https://golang.org/dl/
2. Create a new directory for your project and navigate to it in your terminal.
3. Run `go mod init github.com/yourusername/url-shortener` to initialize a new Go module.
4. Copy the provided code into a file named `main.go` in your project directory.

## Usage

1. Replace `"YOUR_BITLY_API_KEY"` in the `generateShortURL` function with your actual Bitly API key. You can obtain one from https://bitly.com/a/oauth_apps.
2. Build the executable by running `go build`.
3. Run the executable by executing `./url-shortener` (or `url-shortener.exe` on Windows).
4. Enter a long URL when prompted.
5. The program will validate the URL and generate a shortened URL 

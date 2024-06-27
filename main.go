package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// main function is the entry point of the program.
func main() {
	inputURL := getURL("")
	fmt.Println("Original URL:", inputURL)
	fmt.Println("URL Shortener")

	shortURL := generateShortURL(inputURL)
	fmt.Println("Shortened URL:", shortURL)
}

// getURL function prompts the user to enter a URL and validates it.
// It returns the validated URL.
func getURL(inputURL string) string {
	var validURL = regexp.MustCompile(`^(https?:\/\/)?[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?$`)
	for inputURL == "" {
		fmt.Println("Please enter a URL")
		fmt.Scan(&inputURL)

		if !validURL.MatchString(inputURL) {
			fmt.Println("Invalid URL. Please enter a valid URL")
			inputURL = ""
		}
	}
	return inputURL
}

// generateShortURL function takes a long URL as input, makes a request to the Bitly API,
// and returns the shortened URL.
func generateShortURL(longURL string) string {
	bitlyAPIKey := "YOUR_BITLY_API_KEY"
	bitlyAPIEndpoint := "https://api-ssl.bitly.com/v4/shorten"

	// Create JSON payload
	payload := map[string]string{
		"long_url": longURL,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON", err)
		return ""
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", bitlyAPIEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Add("Authorization", "Bearer "+bitlyAPIKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body", err)
		return ""
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response from Bitly:", string(body))
		return ""
	}

	// Parse the response JSON
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error unmarshalling response JSON", err)
		return ""
	}

	shortenedURL, ok := result["link"].(string)
	if !ok {
		fmt.Println("Error retrieving shortened URL from response")
		return ""
	}

	return shortenedURL
}

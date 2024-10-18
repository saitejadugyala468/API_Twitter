package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Function to delete a tweet by ID
func deleteTweet(httpClient *http.Client, tweetID string) {
	apiURL := "https://api.twitter.com/1.1/statuses/destroy/" + tweetID + ".json"

	// Create a new POST request to delete the tweet
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(""))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Error deleting tweet: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Tweet deleted successfully!")
	} else {
		fmt.Printf("Failed to delete tweet. Status: %s\n", resp.Status)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// Function to post a new tweet and return its ID
func postTweet(httpClient *http.Client, tweet string) string {
	apiURL := "https://api.twitter.com/1.1/statuses/update.json"
	data := url.Values{}
	data.Set("status", tweet)

	// Make the POST request to post a new tweet
	resp, err := httpClient.PostForm(apiURL, data)
	if err != nil {
		log.Fatalf("Error posting tweet: %v", err)
	}
	defer resp.Body.Close()

	// Check if tweet was posted successfully
	if resp.StatusCode == http.StatusOK {
		// Parse the JSON response to extract tweet ID
		var responseMap map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&responseMap)
		tweetID := responseMap["id_str"].(string)
		fmt.Println("Tweet posted successfully! Tweet ID:", tweetID)
		return tweetID
	} else {
		fmt.Printf("Failed to post tweet. Status: %s\n", resp.Status)
		return ""
	}
}

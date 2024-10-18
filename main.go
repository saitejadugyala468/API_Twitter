package main

import (
	"fmt"

	"github.com/dghubble/oauth1"
)

// Twitter API credentials (replace with your keys)
const (
	apiKey            = "Lsl9K7XWlt1XPnMt3DqPWhl7B"
	apiSecretKey      = "oagZjdct1z8S31WWEOUt8NMNtrH6cTIRuBpjCOoC01qZZUMTHA"
	accessToken       = "1563211586982256645-tWX7TnPDB67AEFoxaZdGGvu7TRw1SS"
	accessTokenSecret = "FltfmRCp8ivU0KkBPX7tioZ1p6qzMSHzkksT3GLQv9dul"
)

func main() {
	// Initialize OAuth1a configuration
	config := oauth1.NewConfig(apiKey, apiSecretKey)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Post a tweet and capture the tweet ID
	fmt.Println("Attempting to post a tweet...")
	tweetID := postTweet(httpClient, "Hello from Twitter API using Go!")

	if tweetID != "" {
		fmt.Println("Posted tweet with ID:", tweetID)

		// Delete the tweet using the captured tweet ID
		fmt.Println("Attempting to delete the tweet...")
		deleteTweet(httpClient, tweetID)
	} else {
		fmt.Println("Failed to post tweet, skipping deletion.")
	}
}

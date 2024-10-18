# API_Twitter
Setup Instructions
# Twitter Developer Account Setup
We have to cleate a Twitter Developer Account, create one by visiting the Twitter Developer Platform.
Apply for a developer account, and once approved:
Create a new project and an app.
Set the app permissions to Read and Write under User Authentication Settings.
Set the Callback URL to http://localhost:3000 (for local testing purposes).
Generate API Keys and Tokens
Navigate to Projects & Apps → Your App → Keys and Tokens.
Generate the following keys and tokens:
API Key
API Secret Key
Bearer Token (if using OAuth 2.0, but we are using OAuth 1.0a in this case)
Access Token
Access Token Secret
Copy and save these keys securely. They will be used for authentication in the program.
Go Installation
Install Go by following the instructions at Go's official site.
Verify the installation with:
bash
Copy code
go version
Project Setup
Clone the repository or download the project files.
Navigate to the project folder:
bash
Copy code
cd Twitter_API
Open the main.go file and replace the placeholders with your actual Twitter API credentials:
go
Copy code
const (
    apiKey            = "your-api-key"
    apiSecretKey      = "your-api-secret-key"
    accessToken       = "your-access-token"
    accessTokenSecret = "your-access-token-secret"
)
Run the program using:
bash
Copy code
go run main.go
Program Details
Post a Tweet
The program uses the statuses/update.json endpoint to post a new tweet. You can modify the content of the tweet in the postTweet function.

# Example:

go
Copy code
postTweet(httpClient, "Hello from Twitter API using Go!")
Upon successfully posting, the tweet ID is returned and printed to the console.

# Delete a Tweet
The program uses the statuses/destroy/:id.json endpoint to delete a tweet. The tweet ID is required as a parameter. The tweet ID can be captured after posting or retrieved manually from an existing tweet.

# Example:

go
Copy code
deleteTweet(httpClient, tweetID)
Error Handling
The program implements basic error handling for common issues such as:

# Invalid credentials:
If any of the provided OAuth credentials are incorrect, the Twitter API will return a 401 Unauthorized status.
Rate limiting: If too many requests are made in a short period, the API will return a 429 status indicating rate limiting.
Invalid tweet ID: If an invalid tweet ID is provided for deletion, the API will return a 404 Not Found status.
The program prints meaningful messages in case of such errors.

# Example API Requests and Responses
Post a Tweet Request
Endpoint: POST https://api.twitter.com/1.1/statuses/update.json
## Parameters:
status: The content of the tweet (e.g., "Hello from Twitter API using Go!")
# Response Example:
json
Copy code
{
  "created_at": "Wed Oct 17 20:00:00 +0000 2024",
  "id": 1234567890123456789,
  "id_str": "1234567890123456789",
  "text": "Hello from Twitter API using Go!",
  ...
}
## Delete a Tweet Request
Endpoint: POST https://api.twitter.com/1.1/statuses/destroy/:id.json
Parameters:
id: The ID of the tweet to be deleted (e.g., 1234567890123456789)
## Response Example:
json
Copy code
{
  "truncated": false,
  "id": 1234567890123456789,
  "id_str": "1234567890123456789",
  "text": "Hello from Twitter API using Go!"
}
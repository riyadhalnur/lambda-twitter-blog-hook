package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	hookSecret     = os.Getenv("HOOK_SECRET")
	consumerKey    = os.Getenv("API_KEY")
	consumerSecret = os.Getenv("API_SECRET")
	accessToken    = os.Getenv("ACCESS_TOKEN")
	accessSecret   = os.Getenv("ACCESS_SECRET")
	baseLink       = os.Getenv("BASE_URL")
	comTitle       = "[NEW POST]"
)

// Handler - handles the Lambda event
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Check if webhook secret header exists and value is equal to defined secret
	if hookHeader := request.Headers["X-Gitlab-Token"]; hookHeader != hookSecret {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    getCorsHeaders(),
			Body:       "invalid webhook secret",
		}, nil
	}

	var bodyObj GitlabPushEvent
	if err := json.Unmarshal([]byte(request.Body), &bodyObj); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers:    getCorsHeaders(),
			Body:       "invalid push event",
		}, err
	}

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)

	commitMessage := strings.Split(bodyObj.Commits[0].Message, "\n")
	if len(commitMessage) < 3 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers:    getCorsHeaders(),
			Body:       "does not seem like a new post. skipping",
		}, nil
	}

	commitTitle, postName, postLink := commitMessage[0], commitMessage[1], commitMessage[2]
	if commitTitle != comTitle {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers:    getCorsHeaders(),
			Body:       "not a new post. skipping",
		}, nil
	}

	newTweet := fmt.Sprintf("New post 🗞️ 🗞️ - %s %s%s", postName, baseLink, postLink)
	tweet, err := api.PostTweet(newTweet, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    getCorsHeaders(),
			Body:       "failed to post tweet",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    getCorsHeaders(),
		Body:       tweet.Text,
	}, nil
}

func getCorsHeaders() map[string]string {
	return map[string]string{
		"Access-Control-Allow-Origin": "*",
	}
}

func main() {
	lambda.Start(Handler)
}

package main

import (
	"encoding/json"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type errorResponse struct {
	Reason string `json:"reason"`
}

type successResponse struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

var (
	hookSecret     = os.Getenv("HOOK_SECRET")
	consumerKey    = os.Getenv("API_KEY")
	consumerSecret = os.Getenv("API_SECRET")
	accessToken    = os.Getenv("ACCESS_TOKEN")
	accessSecret   = os.Getenv("ACCESS_SECRET")
)

// Handler - handles the Lambda event
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Check if webhook secret header exists and value is equal to defined secret
	if hookHeader := request.Headers["X-Gitlab-Token"]; hookHeader != hookSecret {
		errBody, err := json.Marshal(&errorResponse{Reason: "Invalid webhook"})
		if err != nil {
			return events.APIGatewayProxyResponse{Body: "internal server error", StatusCode: 500}, err
		}
		return events.APIGatewayProxyResponse{Body: string(errBody), StatusCode: 403}, nil
	}

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessSecret)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}

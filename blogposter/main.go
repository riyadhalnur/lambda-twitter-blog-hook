package main

import (
	"encoding/json"
	"os"
	"time"

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

type pushEvent struct {
	ObjectKind   string      `json:"object_kind"`
	EventName    string      `json:"event_name"`
	Before       string      `json:"before"`
	After        string      `json:"after"`
	Ref          string      `json:"ref"`
	CheckoutSha  string      `json:"checkout_sha"`
	Message      interface{} `json:"message"`
	UserID       int         `json:"user_id"`
	UserName     string      `json:"user_name"`
	UserUsername string      `json:"user_username"`
	UserEmail    string      `json:"user_email"`
	UserAvatar   string      `json:"user_avatar"`
	ProjectID    int         `json:"project_id"`
	Project      struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebURL            string      `json:"web_url"`
		AvatarURL         interface{} `json:"avatar_url"`
		GitSSHURL         string      `json:"git_ssh_url"`
		GitHTTPURL        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      interface{} `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		URL               string      `json:"url"`
		SSHURL            string      `json:"ssh_url"`
		HTTPURL           string      `json:"http_url"`
	} `json:"project"`
	Commits []struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
		Added    []string      `json:"added"`
		Modified []interface{} `json:"modified"`
		Removed  []interface{} `json:"removed"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
	Repository        struct {
		Name            string `json:"name"`
		URL             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
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

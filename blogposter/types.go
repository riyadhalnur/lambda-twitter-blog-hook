package main

import "time"

// GitlabPushEvent struct
type GitlabPushEvent struct {
	ObjectKind   string      `json:"object_kind,omitempty"`
	EventName    string      `json:"event_name,omitempty"`
	Before       string      `json:"before,omitempty"`
	After        string      `json:"after,omitempty"`
	Ref          string      `json:"ref,omitempty"`
	CheckoutSha  string      `json:"checkout_sha,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	UserID       int         `json:"user_id,omitempty"`
	UserName     string      `json:"user_name,omitempty"`
	UserUsername string      `json:"user_username,omitempty"`
	UserEmail    string      `json:"user_email,omitempty"`
	UserAvatar   string      `json:"user_avatar,omitempty"`
	ProjectID    int         `json:"project_id,omitempty"`
	Project      struct {
		ID                int         `json:"id,omitempty"`
		Name              string      `json:"name,omitempty"`
		Description       string      `json:"description,omitempty"`
		WebURL            string      `json:"web_url,omitempty"`
		AvatarURL         interface{} `json:"avatar_url,omitempty"`
		GitSSHURL         string      `json:"git_ssh_url,omitempty"`
		GitHTTPURL        string      `json:"git_http_url,omitempty"`
		Namespace         string      `json:"namespace,omitempty"`
		VisibilityLevel   int         `json:"visibility_level,omitempty"`
		PathWithNamespace string      `json:"path_with_namespace,omitempty"`
		DefaultBranch     string      `json:"default_branch,omitempty"`
		CiConfigPath      interface{} `json:"ci_config_path,omitempty"`
		Homepage          string      `json:"homepage,omitempty"`
		URL               string      `json:"url,omitempty"`
		SSHURL            string      `json:"ssh_url,omitempty"`
		HTTPURL           string      `json:"http_url,omitempty"`
	} `json:"project,omitempty"`
	Commits []struct {
		ID        string    `json:"id,omitempty"`
		Message   string    `json:"message,omitempty"`
		Timestamp time.Time `json:"timestamp,omitempty"`
		URL       string    `json:"url,omitempty"`
		Author    struct {
			Name  string `json:"name,omitempty"`
			Email string `json:"email,omitempty"`
		} `json:"author,omitempty"`
		Added    []string      `json:"added,omitempty"`
		Modified []interface{} `json:"modified,omitempty"`
		Removed  []interface{} `json:"removed,omitempty"`
	} `json:"commits,omitempty"`
	TotalCommitsCount int `json:"total_commits_count,omitempty"`
	Repository        struct {
		Name            string `json:"name,omitempty"`
		URL             string `json:"url,omitempty"`
		Description     string `json:"description,omitempty"`
		Homepage        string `json:"homepage,omitempty"`
		GitHTTPURL      string `json:"git_http_url,omitempty"`
		GitSSHURL       string `json:"git_ssh_url,omitempty"`
		VisibilityLevel int    `json:"visibility_level,omitempty"`
	} `json:"repository,omitempty"`
}

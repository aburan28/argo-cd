package pull_request

import (
	"github.com/argoproj/argo-cd/v2/applicationset/services/github_app_auth"
	"github.com/argoproj/argo-cd/v2/applicationset/services/internal/github_app"
)

func NewGithubAppService(g github_app_auth.Authentication, url, owner, repo string, labels []string, cacheEnabled bool) (PullRequestService, error) {
	client, err := github_app.Client(g, url, cacheEnabled)
	if err != nil {
		return nil, err
	}
	return &GithubService{
		client: client,
		owner:  owner,
		repo:   repo,
		labels: labels,
	}, nil
}

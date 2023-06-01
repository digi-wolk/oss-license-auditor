package github

import (
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"testing"
)

// Test FetchGithubRepoDetails updates license for github.com/gin-gonic/gin
func TestFetchGithubRepoDetails(t *testing.T) {
	expectedLicenseId := "MIT"
	githubRepo := &types.GithubRepo{
		Owner: "gin-gonic",
		Repo:  "gin",
	}
	err := FetchGithubRepoDetails(githubRepo)
	if err != nil {
		t.Error("FetchGithubRepoDetails failed")
	}
	if githubRepo.LicenseId != expectedLicenseId {
		t.Error("FetchGithubRepoDetails was incorrect, got: " + githubRepo.LicenseId + ", want: " + expectedLicenseId)
	}
}

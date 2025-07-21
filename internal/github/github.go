package github

import (
	"encoding/json"
	"fmt"
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"log"
	"net/http"
	"os"
)

// GitHubRepoFetcher is a function type for fetching GitHub repo details
type GitHubRepoFetcher func(githubRepo *types.GithubRepo) error

// currentFetcher is the function that will be used to fetch GitHub repo details
// By default, it's set to the real implementation
var currentFetcher GitHubRepoFetcher = fetchGithubRepoDetailsReal

// FetchGithubRepoDetails updates Github repo information
// It uses the currentFetcher which can be replaced with a mock for testing
func FetchGithubRepoDetails(githubRepo *types.GithubRepo) error {
	return currentFetcher(githubRepo)
}

// SetGitHubRepoFetcher sets the function that will be used to fetch GitHub repo details
// This can be used to replace the real implementation with a mock for testing
func SetGitHubRepoFetcher(fetcher GitHubRepoFetcher) {
	currentFetcher = fetcher
}

// ResetGitHubRepoFetcher resets the fetcher to the real implementation
func ResetGitHubRepoFetcher() {
	currentFetcher = fetchGithubRepoDetailsReal
}

// fetchGithubRepoDetailsReal is the real implementation of FetchGithubRepoDetails
func fetchGithubRepoDetailsReal(githubRepo *types.GithubRepo) error {
	githubToken := getGithubToken()
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", githubRepo.Owner, githubRepo.Repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if githubToken != "" {
		req.Header.Set("Authorization", "Bearer "+githubToken)
		req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}

	if resp.StatusCode != 200 {
		log.Fatalf("error calling Github API. Returned status code: %d and status message: %s", resp.StatusCode, resp.Status)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}
	licenseId := result["license"].(map[string]interface{})["spdx_id"].(string)
	githubRepo.LicenseId = licenseId

	return nil
}

func getGithubToken() string {
	token, exists := os.LookupEnv("GITHUB_TOKEN")
	if exists {
		return token
	}
	return ""
}

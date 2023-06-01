package github

import (
	"encoding/json"
	"fmt"
	"github.com/digi-wolk/oss-license-auditor/types"
	"log"
	"net/http"
	"os"
)

// FetchGithubRepoDetails updates Github repo information
func FetchGithubRepoDetails(githubRepo *types.GithubRepo) error {
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

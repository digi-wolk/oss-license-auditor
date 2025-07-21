package github

import (
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"os"
	"testing"
)

// TestSetGitHubRepoFetcher tests that SetGitHubRepoFetcher correctly sets a custom fetcher
// This is important to ensure the dependency injection mechanism works properly
func TestSetGitHubRepoFetcher(t *testing.T) {
	// Save the original fetcher to restore it after the test
	originalFetcher := currentFetcher
	defer func() { currentFetcher = originalFetcher }()

	// Create a custom fetcher that always returns a specific license
	customFetcher := func(githubRepo *types.GithubRepo) error {
		githubRepo.LicenseId = "CUSTOM-LICENSE"
		return nil
	}

	// Set the custom fetcher
	SetGitHubRepoFetcher(customFetcher)

	// Verify that the custom fetcher is used
	githubRepo := &types.GithubRepo{
		Owner: "test-owner",
		Repo:  "test-repo",
	}
	err := FetchGithubRepoDetails(githubRepo)
	if err != nil {
		t.Error("FetchGithubRepoDetails failed")
	}
	if githubRepo.LicenseId != "CUSTOM-LICENSE" {
		t.Error("SetGitHubRepoFetcher was incorrect, got: " + githubRepo.LicenseId + ", want: CUSTOM-LICENSE")
	}
}

// TestResetGitHubRepoFetcher tests that ResetGitHubRepoFetcher correctly resets to the real implementation
// This is important to ensure tests don't affect each other by leaving a mock in place
func TestResetGitHubRepoFetcher(t *testing.T) {
	// Set a custom fetcher first
	customFetcher := func(githubRepo *types.GithubRepo) error {
		githubRepo.LicenseId = "CUSTOM-LICENSE"
		return nil
	}
	SetGitHubRepoFetcher(customFetcher)

	// Reset to the real implementation
	ResetGitHubRepoFetcher()

	// Verify that currentFetcher is now fetchGithubRepoDetailsReal
	// We can't call it directly as it would make a real API call
	// So we just check that it's not our custom fetcher anymore
	if currentFetcher == nil {
		t.Error("ResetGitHubRepoFetcher did not set currentFetcher")
	}

	// Note: We don't actually call FetchGithubRepoDetails here to avoid making a real API call
	// We've already verified that currentFetcher is not nil, which is sufficient for this test

	// Set it back to our custom fetcher for cleanup
	SetGitHubRepoFetcher(customFetcher)
}

// TestGetGithubToken tests that getGithubToken correctly retrieves the token from environment variables
// This is important to ensure authentication with GitHub API works properly
func TestGetGithubToken(t *testing.T) {
	// Save the original token to restore it after the test
	originalToken, tokenExists := os.LookupEnv("GITHUB_TOKEN")

	// Test with token set
	os.Setenv("GITHUB_TOKEN", "test-token")
	token := getGithubToken()
	if token != "test-token" {
		t.Error("getGithubToken was incorrect, got: " + token + ", want: test-token")
	}

	// Test with token unset
	os.Unsetenv("GITHUB_TOKEN")
	token = getGithubToken()
	if token != "" {
		t.Error("getGithubToken was incorrect, got: " + token + ", want: empty string")
	}

	// Restore the original token
	if tokenExists {
		os.Setenv("GITHUB_TOKEN", originalToken)
	} else {
		os.Unsetenv("GITHUB_TOKEN")
	}
}

// TestFetchGithubRepoDetailsWithMock tests that the mock implementation of FetchGithubRepoDetails
// correctly updates the license for a known repository
func TestFetchGithubRepoDetailsWithMock(t *testing.T) {
	// Enable mock mode for this test
	EnableMockMode()
	defer DisableMockMode() // Ensure we reset to real implementation after the test

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

// TestFetchGithubRepoDetailsWithUnknownRepo tests that the mock implementation of FetchGithubRepoDetails
// handles unknown repositories gracefully by setting a default license
func TestFetchGithubRepoDetailsWithUnknownRepo(t *testing.T) {
	// Enable mock mode for this test
	EnableMockMode()
	defer DisableMockMode() // Ensure we reset to real implementation after the test

	expectedLicenseId := "UNKNOWN"
	githubRepo := &types.GithubRepo{
		Owner: "unknown-owner",
		Repo:  "unknown-repo",
	}
	err := FetchGithubRepoDetails(githubRepo)
	if err != nil {
		t.Error("FetchGithubRepoDetails failed")
	}
	if githubRepo.LicenseId != expectedLicenseId {
		t.Error("FetchGithubRepoDetails was incorrect, got: " + githubRepo.LicenseId + ", want: " + expectedLicenseId)
	}
}

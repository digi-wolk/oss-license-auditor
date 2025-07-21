package github

import (
	"github.com/digi-wolk/oss-license-auditor/internal/types"
)

// MockFetchGithubRepoDetails is a mock implementation of FetchGithubRepoDetails
// It can be used in tests to avoid making real API calls to GitHub
func MockFetchGithubRepoDetails(githubRepo *types.GithubRepo) error {
	// Map of known repositories and their licenses
	knownRepos := map[string]map[string]string{
		"gin-gonic": {
			"gin": "MIT",
		},
		"ugorji": {
			"go": "MIT",
		},
		"x": {
			"crypto": "BSD-3-Clause",
		},
	}

	// Check if the owner exists in our known repos
	if ownerRepos, ok := knownRepos[githubRepo.Owner]; ok {
		// Check if the repo exists for this owner
		if license, ok := ownerRepos[githubRepo.Repo]; ok {
			githubRepo.LicenseId = license
			return nil
		}
	}

	// Default license for unknown repos
	githubRepo.LicenseId = "UNKNOWN"
	return nil
}

// EnableMockMode enables mock mode for GitHub API calls
func EnableMockMode() {
	SetGitHubRepoFetcher(MockFetchGithubRepoDetails)
}

// DisableMockMode disables mock mode for GitHub API calls
func DisableMockMode() {
	ResetGitHubRepoFetcher()
}

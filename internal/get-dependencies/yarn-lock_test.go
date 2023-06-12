package get_dependencies

import (
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"testing"
)

// Test GetDependenciesPnpmLock with a valid pnpm-lock.yml file
func TestGetDependenciesYarnLock(t *testing.T) {
	expectedPackages := []struct {
		Name    string
		Version string
	}{
		{Name: "axios", Version: "1.4.0"},
		{Name: "combined-stream", Version: "1.0.8"},
	}

	path := "../../test/fixtures/get-dependencies/yarn-lock/yarn.lock"
	var dependencies types.Dependencies
	dependencies.PackageManagerFile = path
	err := GetDependenciesYarnLock(&dependencies)
	if err != nil {
		t.Errorf("GetDependenciesYarnLock errored: %s", err)
	}

	// Check if all expected packages are present
	for expectedPackage := range expectedPackages {
		found := false
		for _, pkg := range dependencies.Packages {
			if pkg.Name == expectedPackages[expectedPackage].Name && pkg.Version == expectedPackages[expectedPackage].Version {
				found = true
			}
		}
		if !found {
			t.Errorf("GetDependenciesYarnLock was incorrect, expected %s package not found.", expectedPackages[expectedPackage].Name)
		}
	}
}

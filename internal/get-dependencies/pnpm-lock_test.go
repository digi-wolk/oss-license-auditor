package get_dependencies

import (
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"testing"
)

// Test GetDependenciesPnpmLock with a valid pnpm-lock.yml file
func TestGetDependenciesPnpmLock(t *testing.T) {
	expectedPackages := []struct {
		Name    string
		Version string
	}{
		{Name: "lodash", Version: "4.17.21"},
		{Name: "@ampproject/remapping", Version: "2.2.1"},
		{Name: "react-dom", Version: "18.2.0"},
	}

	path := "../../test/fixtures/get-dependencies/pnpm-lock/pnpm-lock.yaml"
	var dependencies types.Dependencies
	dependencies.PackageManagerFile = path
	err := GetDependenciesPnpmLock(&dependencies)
	if err != nil {
		t.Errorf("GetDependenciesPackageJsonLock errored: %s", err)
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
			t.Errorf("GetDependenciesPnpmLock was incorrect, expected %s package not found.", expectedPackages[expectedPackage].Name)
		}
	}
}

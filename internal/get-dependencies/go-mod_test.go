package get_dependencies

import (
	"github.com/digi-wolk/oss-license-auditor/internal/github"
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"log"
	"testing"
)

// TestGetDependenciesWithTestProjectPath tests that GetDependenciesGoMod correctly parses
// dependencies from a go.mod file and finds the expected package
func TestGetDependenciesWithTestProjectPath(t *testing.T) {
	// Enable mock mode for GitHub API calls to avoid real API calls during tests
	github.EnableMockMode()
	defer github.DisableMockMode() // Ensure we reset to real implementation after the test

	projectPath := "../../test/fixtures/get-dependencies/go-go-mod/go.mod"
	var dependencies types.Dependencies
	dependencies.PackageManagerFile = projectPath
	err := GetDependenciesGoMod(&dependencies)
	if err != nil {
		t.Errorf("ListPackages was incorrect, got: %s, want: nil.", err)
	}
	expectedPackage := types.Package{
		Name:    "gin",
		Owner:   "gin-gonic",
		Version: "v1.7.5",
	}
	// Loop through packages of type GoModPackage
	for _, pkg := range dependencies.Packages {
		log.Printf("pkg.Name: %s pkg.Owner: %s pkg.Version: %s", pkg.Name, pkg.Owner, pkg.Version)
		// If the package name and version match the expected package name and version, return
		if pkg.Name == expectedPackage.Name && pkg.Owner == expectedPackage.Owner {
			return
		}
	}
	t.Errorf("ListPackages was incorrect, expected %s package %s version not found.", expectedPackage.Name, expectedPackage.Version)
}

// TestGetDependenciesWithLicenseInfo tests that GetDependenciesGoMod correctly fetches
// license information for dependencies using the GitHub API (mocked in this test)
func TestGetDependenciesWithLicenseInfo(t *testing.T) {
	// Enable mock mode for GitHub API calls to avoid real API calls during tests
	github.EnableMockMode()
	defer github.DisableMockMode() // Ensure we reset to real implementation after the test

	projectPath := "../../test/fixtures/get-dependencies/go-go-mod/go.mod"
	var dependencies types.Dependencies
	dependencies.PackageManagerFile = projectPath
	err := GetDependenciesGoMod(&dependencies)
	if err != nil {
		t.Errorf("GetDependenciesGoMod failed with error: %s", err)
	}

	// Check that we have at least one package with license information
	foundLicense := false
	for _, pkg := range dependencies.Packages {
		if pkg.License != "" && pkg.License != "UNKNOWN (empty)" {
			foundLicense = true
			break
		}
	}

	if !foundLicense {
		t.Errorf("GetDependenciesGoMod did not fetch license information for any package")
	}
}

// Test GetPackageInfoFromGoModLine returns the correct package info
func TestGetPackageInfoFromGoModLine3Parts(t *testing.T) {
	// Example go.mod file line
	// github.com/ugorji/go/codec v1.1.7 // indirect
	// golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	goModFileLine := "github.com/ugorji/go/codec v1.1.7 // indirect"
	expectedGoModPackageInfo := GoModPackageInfo{
		Provider: "github.com",
		FullName: "ugorji/go",
		Owner:    "ugorji",
		Name:     "go",
		Version:  "v1.1.7",
	}
	goModPackageInfo := GetPackageInfoFromGoModLine(goModFileLine)
	if goModPackageInfo != expectedGoModPackageInfo {
		t.Errorf("GetPackageInfoFromGoModLine was incorrect, got: %s, want: %s.", goModPackageInfo, expectedGoModPackageInfo)
	}
}

// Test GetPackageInfoFromGoModLine returns the correct package info
func TestGetPackageInfoFromGoModLine2Parts(t *testing.T) {
	// Example go.mod file line
	// golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	goModFileLine := "golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect"
	expectedGoModPackageInfo := GoModPackageInfo{
		Provider: "golang.org",
		FullName: "x/crypto",
		Owner:    "x",
		Name:     "crypto",
		Version:  "v0.0.0-20200622213623-75b288015ac9",
	}
	goModPackageInfo := GetPackageInfoFromGoModLine(goModFileLine)
	if goModPackageInfo != expectedGoModPackageInfo {
		t.Errorf("GetPackageInfoFromGoModLine was incorrect, got: %s, want: %s.", goModPackageInfo, expectedGoModPackageInfo)
	}
}

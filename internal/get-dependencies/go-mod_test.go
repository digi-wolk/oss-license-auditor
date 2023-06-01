package get_dependencies

import (
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"log"
	"testing"
)

// Test GetDependenciesGoMod contains github.com/gin-gonic/gin v1.7.5
func TestGetDependenciesWithTestProjectPath(t *testing.T) {
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
		// Print length of packages
		// If the package name and version match the expected package name and version, return
		if pkg.Name == expectedPackage.Name && pkg.Owner == expectedPackage.Owner {
			return
		}
	}
	t.Errorf("ListPackages was incorrect, expected %s package %s version not found.", expectedPackage.Name, expectedPackage.Version)
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

package npm

import (
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"testing"
)

// Test UpdatePackageFromNpm updates NPM license for @babel/code-frame
func TestUpdatePackageFromNpmBabel(t *testing.T) {
	expectedLicense := "MIT"
	npmPackage := types.Package{
		Owner:   "@babel",
		Name:    "code-frame",
		Version: "7.12.13",
	}
	err := UpdatePackageFromNpm(&npmPackage)
	if err != nil {
		t.Error("UpdatePackageFromNpm failed")
	}
	if npmPackage.License != expectedLicense {
		t.Error("UpdatePackageFromNpm was incorrect, got: " + npmPackage.License + ", want: " + expectedLicense)
	}
}

// Test UpdatePackageFromNpm updates NPM license for @babel/code-frame
func TestUpdatePackageFromNpmPrelude(t *testing.T) {
	expectedLicense := "MIT"
	npmPackage := types.Package{
		Owner:   "",
		Name:    "prelude-ls",
		Version: "1.1.2",
	}
	err := UpdatePackageFromNpm(&npmPackage)
	if err != nil {
		t.Error("UpdatePackageFromNpm failed")
	}
	if npmPackage.License != expectedLicense {
		t.Error("UpdatePackageFromNpm was incorrect, got: " + npmPackage.License + ", want: " + expectedLicense)
	}
}

// Test UpdatePackageFromNpm updates NPM license for react-loadable (version 5.5.2) not found on NPM
func TestUpdatePackageFromNpmReactLoadable(t *testing.T) {
	expectedLicense := "UNKNOWN (not found)"
	npmPackage := types.Package{
		Owner:   "",
		Name:    "react-loadable",
		Version: "5.5.2",
	}
	err := UpdatePackageFromNpm(&npmPackage)
	if err != nil {
		t.Error("UpdatePackageFromNpm failed")
	}
	if npmPackage.License != expectedLicense {
		t.Error("UpdatePackageFromNpm was incorrect, got: " + npmPackage.License + ", want: " + expectedLicense)
	}
}

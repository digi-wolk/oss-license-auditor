package npm

import (
	"github.com/digi-wolk/oss-license-auditor/types"
	"testing"
)

// Test UpdatePackageFromNpm updates NPM license for @babel/code-frame
func TestUpdatePackageFromNpm(t *testing.T) {
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

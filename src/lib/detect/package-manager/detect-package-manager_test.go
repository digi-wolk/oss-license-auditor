package lang

import (
	"testing"
)

// Test DetectUsedPackageManagers with package.json and go.mod
func TestDetectUsedPackageManagers(t *testing.T) {
	fixturePath := "../../../../test-fixtures/detect-package-managers"
	detectedPackageManagers := DetectUsedPackageManagers(fixturePath)
	expectedNumberOfDetectedPackageManagers := 2
	if len(detectedPackageManagers) != expectedNumberOfDetectedPackageManagers {
		t.Errorf("DetectUsedPackageManagers was incorrect, got: %d, want: %d.", len(detectedPackageManagers), expectedNumberOfDetectedPackageManagers)
	}
}

// Test DetectUsedPackageManagers with can detect go.mod
func TestDetectUsedPackageManagersCanDetectGoMod(t *testing.T) {
	fixturePath := "../../../../test-fixtures/detect-package-managers"
	detectedPackageManagers := DetectUsedPackageManagers(fixturePath)
	if detectedPackageManagers[fixturePath+"/src-1/go.mod"] != "go" {
		t.Errorf("DetectUsedPackageManagers was incorrect, got: %s, want: %s.", detectedPackageManagers["go.mod"], "go")
	}
}

// Test DetectUsedPackageManagers with can detect package.json
func TestDetectUsedPackageManagersCanDetectPackageJson(t *testing.T) {
	fixturePath := "../../../../test-fixtures/detect-package-managers"
	detectedPackageManagers := DetectUsedPackageManagers(fixturePath)
	if detectedPackageManagers[fixturePath+"/src-2/package-lock.json"] != "npm" {
		t.Errorf("DetectUsedPackageManagers was incorrect, got: %s, want: %s.", detectedPackageManagers["package-lock.json"], "npm")
	}
}

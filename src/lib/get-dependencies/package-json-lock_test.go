package get_dependencies

import (
	"github.com/digi-wolk/oss-license-auditor/types"
	"testing"
)

// Test GetDependenciesPackageJsonLock with a valid package-lock.json file contains expected package
func TestGetDependenciesPackageJsonLockWithAValidPackageLockJsonFile(t *testing.T) {
	expectedPackageOwner := "@humanwhocodes"
	expectedPackageName := "object-schema"
	path := "../../../test-fixtures/get-dependencies/npm-package-lock/package-lock.json"
	var dependencies types.Dependencies
	dependencies.PackageManagerFile = path
	err := GetDependenciesPackageJsonLock(&dependencies)
	if err != nil {
		t.Errorf("GetDependenciesPackageJsonLock errored: %s", err)
	}
	// Loop through packages of type PackageJsonLockPackage
	for _, pkg := range dependencies.Packages {
		if pkg.Owner != "" {
			t.Logf("pkg.Name: %s pkg.Owner: %s pkg.Version: %s", pkg.Name, pkg.Owner, pkg.Version)
		}
		if pkg.Name == expectedPackageName && pkg.Owner == expectedPackageOwner {
			return
		}
	}
	t.Errorf("GetDependenciesPackageJsonLock was incorrect, expected %s package not found.", expectedPackageName)
}

// Test openFileWithRetry with a valid package-lock.json file
func TestOpenFileWithRetryWithAValidPackageLockJsonFile(t *testing.T) {
	path := "../../../test-fixtures/detect-package-managers/src-2/package-lock.json"
	_, err := openFileWithRetry(path)
	if err != nil {
		t.Errorf("openFileWithRetry errored: %s", err)
	}
}

// Test extractNameFromFullName
func TestExtractNameFromFullName(t *testing.T) {
	// For example: inquirer/node_modules/ansi-regex should become ansi-regex
	// For example: node_modules/ansi-regex should become ansi-regex
	// For example: ansi-regex should become ansi-regex
	// For example: @babel/code-frame should become @babel/code-frame
	// For example: node_modules/@types/json5 should become json5
	fromTo := map[string]string{
		"node_modules/@types/json5":        "json5",
		"inquirer/node_modules/ansi-regex": "ansi-regex",
		"node_modules/ansi-regex":          "ansi-regex",
		"ansi-regex":                       "ansi-regex",
		"@babel/code-frame":                "@babel/code-frame",
	}
	for from, to := range fromTo {
		result := extractNameFromFullName(from)
		if result != to {
			t.Errorf("removeNodeModulesPrefix expected %s, got %s", to, result)
		}
	}
}

// Test extractOwnerFromFullName
func TestExtractOwnerFromFullName(t *testing.T) {
	// For example: inquirer/node_modules/ansi-regex should become ""
	// For example: node_modules/ansi-regex should become ""
	// For example: ansi-regex should become ""
	// For example: @babel/code-frame should become @babel
	// For example: node_modules/@types/json5 should become @types
	fromTo := map[string]string{
		"inquirer/node_modules/ansi-regex":        "",
		"node_modules/ansi-regex":                 "",
		"ansi-regex":                              "",
		"inquirer/node_modules/@babel/ansi-regex": "@babel",
		"node_modules/@babel/ansi-regex":          "@babel",
		"@babel/code-frame":                       "@babel",
		"node_modules/@types/json5":               "@types",
	}
	for from, to := range fromTo {
		result := extractOwnerFromFullName(from)
		if result != to {
			t.Errorf("removeNodeModulesPrefix expected %s, got %s", to, result)
		}
	}
}

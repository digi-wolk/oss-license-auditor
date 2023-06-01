package get_dependencies

import (
	"encoding/json"
	"fmt"
	"github.com/digi-wolk/oss-license-auditor/lib/npm"
	"github.com/digi-wolk/oss-license-auditor/types"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// RetryAttempts represents the number of retry attempts
const RetryAttempts = 3

// RetryDelay represents the delay duration between retries
const RetryDelay = time.Second * 2

// GetDependenciesPackageJsonLock Read the package-lock.json file and return the dependencies
func GetDependenciesPackageJsonLock(dependencies *types.Dependencies) error {
	var packageJsonLock PackageJsonLock

	jsonFile, err := openFileWithRetry(dependencies.PackageManagerFile)
	if err != nil {
		log.Fatal("Error opening package-lock.json file: ", err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal("Error closing package-lock.json file: ", err)
		}
	}(jsonFile)
	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &packageJsonLock)
	if err != nil {
		return err
	}

	for packageJsonLockPackageFullName, packageJsonLockPackage := range packageJsonLock.Packages {
		// Skip if package name is empty
		if packageJsonLockPackageFullName == "" {
			continue
		}

		packageInfo := types.Package{
			Owner:   extractOwnerFromFullName(packageJsonLockPackageFullName),
			Name:    extractNameFromFullName(packageJsonLockPackageFullName),
			Version: packageJsonLockPackage.Version,
			Dev:     packageJsonLockPackage.Dev,
		}
		err := npm.UpdatePackageFromNpm(&packageInfo)
		if err != nil {
			return err
		}
		dependencies.Packages = append(dependencies.Packages, packageInfo)
	}
	return nil
}

func openFileWithRetry(filePath string) (*os.File, error) {
	var file *os.File
	var err error

	basePath, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current path: %v", err)
	}
	filePath = filepath.Join(basePath, filepath.Clean(filePath))
	for attempt := 1; attempt <= RetryAttempts; attempt++ {
		file, err = os.Open(filePath)
		if err == nil {
			// File opened successfully, break the retry loop
			return file, nil
		}

		log.Printf("Error opening file %s, retrying in %v...", filePath, RetryDelay)
		time.Sleep(RetryDelay)
	}

	return nil, fmt.Errorf("failed to open file %s after %d attempts: %v", filePath, RetryAttempts, err)
}

// extractNameFromFullName Extract the package name from the full name
func extractNameFromFullName(name string) string {
	nodeModuleIndex := strings.LastIndex(name, "node_modules/")
	if nodeModuleIndex != -1 {
		slashIndex := strings.LastIndex(name, "/")
		if slashIndex == -1 {
			return ""
		}
		return name[slashIndex+1:]
	}
	return name
}

// extractOwnerFromFullName Extract the package owner from the full name
func extractOwnerFromFullName(name string) string {
	nodeModuleIndex := strings.LastIndex(name, "node_modules/")
	if nodeModuleIndex == -1 {
		slashIndex := strings.LastIndex(name, "/")
		if slashIndex == -1 {
			return ""
		}
		return name[:slashIndex]
	}
	// Remove node_modules/ and before it
	name = name[nodeModuleIndex+len("node_modules/"):]
	slashIndex := strings.Index(name, "/")
	if slashIndex == -1 {
		return ""
	}
	return name[:slashIndex]
}

package get_dependencies

import (
	"encoding/json"
	"github.com/digi-wolk/oss-license-auditor/internal/npm"
	"github.com/digi-wolk/oss-license-auditor/internal/shared"
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"io"
	"log"
	"os"
	"strings"
)

// GetDependenciesPackageJsonLock Read the package-lock.json file and return the dependencies
func GetDependenciesPackageJsonLock(dependencies *types.Dependencies) error {
	var packageJsonLock PackageJsonLock

	jsonFile, err := shared.OpenFileWithRetry(dependencies.PackageManagerFile)
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

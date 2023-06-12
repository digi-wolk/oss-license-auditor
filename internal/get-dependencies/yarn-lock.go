package get_dependencies

import (
	"github.com/digi-wolk/oss-license-auditor/internal/npm"
	"github.com/digi-wolk/oss-license-auditor/internal/shared"
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"strings"
)

// PackageValue represents the package information in the yarn.lock file
type PackageValue struct {
	Version      string            `yaml:"version"`
	Resolution   string            `yaml:"resolution"`
	Checksum     string            `yaml:"checksum"`
	LanguageName string            `yaml:"languageName"`
	LinkType     string            `yaml:"linkType"`
	Dependencies map[string]string `yaml:"dependencies"`
}

// GetDependenciesYarnLock reads the yarn.lock file and returns the dependencies
func GetDependenciesYarnLock(dependencies *types.Dependencies) error {
	yamlFile, err := shared.OpenFileWithRetry(dependencies.PackageManagerFile)
	if err != nil {
		log.Fatal("Error opening yarn.lock file: ", err)
	}
	defer func(yamlFile *os.File) {
		err := yamlFile.Close()
		if err != nil {
			log.Fatal("Error closing yarn.lock file: ", err)
		}
	}(yamlFile)
	byteValue, err := io.ReadAll(yamlFile)
	if err != nil {
		log.Fatal("Error reading yarn.lock file: ", err)
	}

	// Parse the YAML data
	var lockfile map[string]PackageValue
	err = yaml.Unmarshal(byteValue, &lockfile)
	if err != nil {
		log.Fatal("Error unmarshaling yarn.lock file: ", err)
	}

	for _, packageValue := range lockfile {
		packageVersion := packageValue.Version
		packageName := strings.Split(packageValue.Resolution, "@")[0]
		packageManager := extractPackageManagerFromFullName(packageValue.Resolution)
		if packageManager != "npm" {
			continue
		}

		packageInfo := types.Package{
			Name:    packageName,
			Version: packageVersion,
		}

		// Update license information from npm
		err = npm.UpdatePackageFromNpm(&packageInfo)
		if err != nil {
			return err
		}

		dependencies.Packages = append(dependencies.Packages, packageInfo)
	}

	return nil
}

func extractPackageManagerFromFullName(resolution string) string {
	atIndex := strings.Index(resolution, "@")
	colonIndex := strings.Index(resolution, ":")
	if atIndex == -1 || colonIndex == -1 {
		return ""
	}

	return resolution[atIndex+1 : colonIndex]
}

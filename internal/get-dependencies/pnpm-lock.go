package get_dependencies

import (
	"github.com/digi-wolk/oss-license-auditor/internal/npm"
	"github.com/digi-wolk/oss-license-auditor/internal/shared"
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type Lockfile struct {
	Dependencies    map[string]Dependency   `yaml:"dependencies"`
	DevDependencies map[string]Dependency   `yaml:"devDependencies"`
	Packages        map[string]PackageEntry `yaml:"packages"`
}

type PackageEntry struct {
	Resolution   Resolution        `yaml:"resolution"`
	Engines      map[string]string `yaml:"engines"`
	Dependencies map[string]string `yaml:"dependencies"`
	Dev          bool              `yaml:"dev"`
}

type Resolution struct {
	Integrity string `yaml:"integrity"`
}

type Dependency struct {
	Specifier string `yaml:"specifier"`
	Version   string `yaml:"version"`
}

// GetDependenciesPnpmLock Read the package-lock.json file and return the dependencies
func GetDependenciesPnpmLock(dependencies *types.Dependencies) error {

	yamlFile, err := shared.OpenFileWithRetry(dependencies.PackageManagerFile)
	if err != nil {
		log.Fatal("Error opening pnpm-lock.yml file: ", err)
	}
	defer func(yamlFile *os.File) {
		err := yamlFile.Close()
		if err != nil {
			log.Fatal("Error closing package-lock.json file: ", err)
		}
	}(yamlFile)
	byteValue, _ := io.ReadAll(yamlFile)

	// Parse the YAML data
	lockfile := Lockfile{}
	err = yaml.Unmarshal(byteValue, &lockfile)
	if err != nil {
		log.Fatal(err)
	}

	for packageKey := range lockfile.Packages {
		// Removes (@babel/core@7.22.5) from /babel-jest@29.5.0(@babel/core@7.22.5)
		unexpectedVersion, err := regexp.Compile(`\(.*?\)`)
		if err == nil {
			packageKey = unexpectedVersion.ReplaceAllString(packageKey, "")
		}

		// Version is anything after the last @
		packageVersion := packageKey[strings.LastIndex(packageKey, "@")+1:]
		// Name is anything before the last @
		packageName := packageKey[:strings.LastIndex(packageKey, "@")]
		// Remove the first character if it is a /
		if packageName[0] == '/' {
			packageName = packageName[1:]
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

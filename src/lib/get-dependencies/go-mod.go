package get_dependencies

import (
	"github.com/digi-wolk/oss-license-auditor/definitions"
	"github.com/digi-wolk/oss-license-auditor/lib/github"
	"github.com/digi-wolk/oss-license-auditor/types"
	"os"
	"os/exec"
	"strings"
)

// GetDependenciesGoMod Get the list of dependencies for a Go project
func GetDependenciesGoMod(dependencies *types.Dependencies) error {

	goModFile := dependencies.PackageManagerFile
	// Get full directory path from the full filename
	projectPath := goModFile[:strings.LastIndex(goModFile, "/")]

	// Get the current OS path
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	// Change the working directory to the project path
	err = os.Chdir(projectPath)
	if err != nil {
		return err
	}

	// Run the "go list -m all" command
	cmd := exec.Command("go", "list", "-m", "all")
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	// os.Chdir back to currentPath
	err = os.Chdir(currentPath)

	// Split the command output into lines
	lines := strings.Split(string(output), "\n")

	// Store the package names in an array
	for _, line := range lines {
		line = strings.TrimSpace(line)
		goModPackageInfo := GetPackageInfoFromGoModLine(line)
		if goModPackageInfo.FullName == "" || goModPackageInfo.Provider == "" || goModPackageInfo.Owner == "" || goModPackageInfo.Name == "" {
			continue
		}
		githubRepo := types.GithubRepo{
			Repo:  goModPackageInfo.Name,
			Owner: goModPackageInfo.Owner,
		}
		if goModPackageInfo.Provider == "github.com" {
			err := github.FetchGithubRepoDetails(&githubRepo)
			if err != nil {
				return err
			}
		}
		// TODO: golang.org
		// TODO: gopkg.in

		packageInfo := types.Package{
			Owner:   goModPackageInfo.Owner,
			Name:    goModPackageInfo.Name,
			Version: goModPackageInfo.Version,
			License: githubRepo.LicenseId,
			Dev:     false,
		}
		if packageInfo.License == "" {
			packageInfo.License = "UNKNOWN"
		} else {
			packageInfo.License = githubRepo.LicenseId
		}
		packageInfo.IsLicenseRiskyFail = definitions.IsLicenseRiskyFail(packageInfo.License)
		packageInfo.IsLicenseRiskyWarn = definitions.IsLicenseRiskyWarn(packageInfo.License)

		dependencies.Packages = append(dependencies.Packages, packageInfo)
	}

	return nil
}

// GetPackageInfoFromGoModLine Get the package info from a line in the go.mod file
func GetPackageInfoFromGoModLine(goModFileLine string) GoModPackageInfo {
	var goModPackageInfo GoModPackageInfo
	index := strings.Index(goModFileLine, "/")
	if index != -1 {
		goModPackageInfo.Provider = goModFileLine[:index]
		goModFileLine = goModFileLine[index+1:]
	}
	index = strings.Index(goModFileLine, " ")
	if index == -1 {
		return goModPackageInfo
	}

	goModPackageInfo.FullName = goModFileLine[:index]
	goModFileLine = goModFileLine[index+1:]
	count := strings.Count(goModPackageInfo.FullName, "/")
	if count == 0 {
		goModPackageInfo.Owner = ""
		goModPackageInfo.Name = goModPackageInfo.FullName
	}
	if count == 1 {
		goModPackageInfo.Owner = goModPackageInfo.FullName[:strings.Index(goModPackageInfo.FullName, "/")]
		goModPackageInfo.Name = goModPackageInfo.FullName[strings.Index(goModPackageInfo.FullName, "/")+1:]
	}
	// Only keep the first two parts like */* if goModPackageInfo.Name is */*/*
	if count == 2 {
		goModPackageInfo.FullName = goModPackageInfo.FullName[:strings.LastIndex(goModPackageInfo.FullName, "/")]
		// First part of packageFullName before /
		goModPackageInfo.Owner = goModPackageInfo.FullName[:strings.Index(goModPackageInfo.FullName, "/")]
		goModPackageInfo.Name = goModPackageInfo.FullName[strings.Index(goModPackageInfo.FullName, "/")+1:]
	}

	// Anything before // if it exists
	index = strings.Index(goModFileLine, "//")
	if index != -1 {
		goModPackageInfo.Version = goModFileLine[:index]
		// Remove whitespace
		goModPackageInfo.Version = strings.TrimSpace(goModPackageInfo.Version)
	} else {
		goModPackageInfo.Version = goModFileLine
	}
	return goModPackageInfo
}

package print_dependencies

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/digi-wolk/oss-license-auditor/internal/cli"
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Comment struct {
	Body string `json:"body"`
}

// ShowDependencies Shows dependencies and returns if it has any risky fail licenses
func ShowDependencies(dependencies *types.Dependencies) bool {
	packages := dependencies.Packages
	packageManagerFile := dependencies.PackageManagerFile
	args := dependencies.CliArguments
	if args.Ci {
		printCi(packageManagerFile, packages, args)
	} else {
		printNonCi(packageManagerFile, packages, args)
	}
	if hasRiskyFailLicenses(packages) {
		return true
	}
	return false
}

func printCi(packageManagerFile string, packages []types.Package, args cli.Arguments) {
	var printPackages []string
	// Show risky fail packages
	for _, packageInfo := range packages {
		if args.Verbose {
			fmt.Println("Package: " + packageInfo.Name + " Version: " + packageInfo.Version + " License: " + packageInfo.License + " Package Manager File: " + packageManagerFile)
		}
		if packageInfo.IsLicenseRiskyFail {
			printPackages = append(printPackages, "| 🛑 | "+getPackageDetailsLine(packageInfo)+" |")
		}
	}
	// Show risky warn packages
	for _, packageInfo := range packages {
		if packageInfo.IsLicenseRiskyWarn {
			printPackages = append(printPackages, "| ⚠️ | "+getPackageDetailsLine(packageInfo)+" |")
		}
	}
	// Show packages with no license determined
	for _, packageInfo := range packages {
		if packageInfo.License == "" || packageInfo.License == "UNKNOWN" {
			printPackages = append(printPackages, "| ⚠️ | "+getPackageDetailsLine(packageInfo)+" |")
		}
	}
	// Show non-risky packages if --only-risky-licenses is not specified
	if !args.OnlyRiskyLicenses {
		for _, packageInfo := range packages {
			if !packageInfo.IsLicenseRiskyFail && !packageInfo.IsLicenseRiskyWarn && packageInfo.License != "" && packageInfo.License != "UNKNOWN" {
				printPackages = append(printPackages, "|  | "+getPackageDetailsLine(packageInfo)+" |")
			}
		}
	}
	if len(printPackages) == 0 {
		log.Println("Nothing to print!")
		return
	}

	var commentMessageLines string

	commentMessageLines += "\n## " + packageManagerFile
	commentMessageLines += "| | Package | Version | License |"
	commentMessageLines += "|-|---------|-------- |---------|"
	for _, printPackage := range printPackages {
		commentMessageLines += printPackage
	}

	if !args.CommentOnGithubPr {
		for commentMessageLine := range commentMessageLines {
			fmt.Println(commentMessageLine)
		}
		return
	}

	owner := os.Getenv("GITHUB_REPOSITORY_OWNER")
	repo := os.Getenv("GITHUB_REPOSITORY")
	pullRequestNumber := os.Getenv("GITHUB_PULL_REQUEST_NUMBER")
	githubToken := os.Getenv("GITHUB_TOKEN")

	err := createPullRequestComment(owner, repo, pullRequestNumber, githubToken, commentMessageLines)
	if err != nil {
		return
	}
}

func printNonCi(packageManagerFile string, packages []types.Package, args cli.Arguments) {
	for _, packageInfo := range packages {
		if args.Verbose {
			fmt.Println("Package: " + packageInfo.Name + " Version: " + packageInfo.Version + " License: " + packageInfo.License + " Package Manager File: " + packageManagerFile)
		}
		// Do not show un-risky licenses if --only-risky-licenses is specified
		if args.OnlyRiskyLicenses && !packageInfo.IsLicenseRiskyFail && !packageInfo.IsLicenseRiskyWarn {
			continue
		}
		fmt.Println(packageInfo)
	}
}

func hasRiskyFailLicenses(packages []types.Package) bool {
	for _, packageInfo := range packages {
		if packageInfo.IsLicenseRiskyFail {
			return true
		}
	}
	return false
}

func getPackageDetailsLine(packageInfo types.Package) string {
	result := fmt.Sprintf("%s | %s | %s", getPackageLink(packageInfo), packageInfo.Version, packageInfo.License)
	return result
}

func getPackageLink(packageInfo types.Package) string {
	fullPackageName := getFullPackageName(packageInfo)
	// Get filename from packageInfo.PackageManagerFile
	PackageManagerFilename := filepath.Base(packageInfo.PackageManagerFile)
	// If it is package-lock.json - link to it
	if PackageManagerFilename == "package-lock.json" {
		return fmt.Sprintf("[%s](https://www.npmjs.com/package/%s)", fullPackageName, fullPackageName)
	}
	return fullPackageName
}

func getFullPackageName(packageInfo types.Package) string {
	if packageInfo.Owner == "" {
		return packageInfo.Name
	}
	return fmt.Sprintf("%s/%s", packageInfo.Owner, packageInfo.Name)
}

func createPullRequestComment(owner string, repo string, pullRequestNumber string, githubToken string, comment string) error {
	if owner == "" {
		return errors.New("github Owner is not set")
	}
	if repo == "" {
		return errors.New("github Repo is not set")
	}
	if pullRequestNumber == "" {
		return errors.New("github Pull Request Number is not set")
	}
	if githubToken == "" {
		return errors.New("github Bearer Token is not set")
	}

	// Construct the comment payload
	commentPayload := Comment{
		Body: comment,
	}
	payloadBytes, err := json.Marshal(commentPayload)
	if err != nil {
		return err
	}

	// Prepare the HTTP request
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s/comments", owner, repo, pullRequestNumber)
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+githubToken)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	// Check the response status
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create comment: %s", resp.Status)
	}

	return nil
}

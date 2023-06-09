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
	"strings"
)

type Comment struct {
	Body string `json:"body"`
}

type GithubActionDetails struct {
	Repository    string
	PullRequestId string
	GithubToken   string
}

func PrintCi(outputLines string, args cli.Arguments) {
	if !args.CommentOnPr {
		for outputLine := range outputLines {
			fmt.Println(outputLine)
		}
		return
	}

	githubActionDetails := GithubActionDetails{
		Repository:    os.Getenv("GITHUB_REPOSITORY"),
		PullRequestId: os.Getenv("INPUT_PULL-REQUEST-ID"),
		GithubToken:   os.Getenv("GITHUB_TOKEN"),
	}

	if args.Verbose {
		log.Println("Github Action Details:")
		log.Println("Github Repository: " + githubActionDetails.Repository)
		log.Println("Pull Request ID: " + githubActionDetails.PullRequestId)
	}

	if args.Verbose {
		log.Println("Will comment on the Pull Request with the following message:")
		log.Println(outputLines)
	}

	err := createPullRequestComment(githubActionDetails.Repository, githubActionDetails.PullRequestId, githubActionDetails.GithubToken, outputLines)
	if err != nil {
		log.Println("Error while creating Pull Request comment: " + err.Error())
		return
	}
}

// GetDependenciesOutput Shows dependencies and returns comment + if it has any risky fail licenses
func GetDependenciesOutput(dependencies *types.Dependencies) string {
	var printPackages []string

	packages := dependencies.Packages
	packageManagerFile := dependencies.PackageManagerFile

	// Show risky fail packages
	for _, packageInfo := range packages {
		if dependencies.CliArguments.Verbose {
			fmt.Println("Package: " + packageInfo.Name + " Version: " + packageInfo.Version + " License: " + packageInfo.License + " Package Manager File: " + packageManagerFile)
		}
		if packageInfo.IsLicenseRiskyFail {
			printPackages = append(printPackages, "\n| 🛑 | "+getPackageDetailsLine(packageInfo)+" |")
		}
	}
	// Show risky warn packages
	for _, packageInfo := range packages {
		if packageInfo.IsLicenseRiskyWarn {
			printPackages = append(printPackages, "\n| ⚠️ | "+getPackageDetailsLine(packageInfo)+" |")
		}
	}
	// Show packages with no license determined
	for _, packageInfo := range packages {
		if packageInfo.License == "" || strings.Contains(packageInfo.License, "UNKNOWN") {
			printPackages = append(printPackages, "\n| ⚠️ | "+getPackageDetailsLine(packageInfo)+" |")
		}
	}
	// Show non-risky packages if --only-risky-licenses is not specified
	if !dependencies.CliArguments.OnlyRiskyLicenses {
		for _, packageInfo := range packages {
			if !packageInfo.IsLicenseRiskyFail && !packageInfo.IsLicenseRiskyWarn && packageInfo.License != "" && !strings.Contains(packageInfo.License, "UNKNOWN") {
				printPackages = append(printPackages, "\n|  | "+getPackageDetailsLine(packageInfo)+" |")
			}
		}
	}
	if len(printPackages) == 0 {
		log.Println("Nothing to print!")
		return "---"
	}

	var commentMessageLines string

	commentMessageLines += "\n\n`" + packageManagerFile + "`"
	commentMessageLines += "\n\n| | Package | Version | License |"
	commentMessageLines += "\n|-|---------|-------- |---------|"
	for _, printPackage := range printPackages {
		commentMessageLines += printPackage
	}
	commentMessageLines += "\n"

	if !dependencies.CliArguments.Ci {
		printDependenciesNonCi(packageManagerFile, packages, dependencies.CliArguments)
	}

	if hasRiskyFailLicenses(packages) {
		dependencies.HasRiskyFailLicense = true
	}
	return commentMessageLines
}

func printDependenciesNonCi(packageManagerFile string, packages []types.Package, args cli.Arguments) {
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

func createPullRequestComment(repository string, pullRequestNumber string, githubToken string, comment string) error {
	if repository == "" {
		return errors.New("github repository is not set")
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
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/issues/%s/comments", repository, pullRequestNumber)
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

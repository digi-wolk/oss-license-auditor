package print_dependencies

import (
	"fmt"
	"github.com/digi-wolk/oss-license-auditor/internal/cli"
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"log"
	"path/filepath"
)

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
	fmt.Println("\n## " + packageManagerFile)
	fmt.Println("| | Package | Version | License |")
	fmt.Println("|-|---------|-------- |---------|")
	for _, printPackage := range printPackages {
		fmt.Println(printPackage)
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

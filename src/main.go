package main

import (
	"fmt"
	"github.com/digi-wolk/oss-license-auditor/lib/cli"
	detectPackageManager "github.com/digi-wolk/oss-license-auditor/lib/detect/package-manager"
	getDependencies "github.com/digi-wolk/oss-license-auditor/lib/get-dependencies"
	printDependencies "github.com/digi-wolk/oss-license-auditor/lib/print-dependencies"
	"github.com/digi-wolk/oss-license-auditor/types"
	"log"
)

func main() {
	var hasRiskyFailLicense bool
	var dependencies types.Dependencies

	args := cli.GetCliArguments()

	// TODO: In the output, alongside the filename, npm or package-lock.json or both?
	usedPackageManagers := detectPackageManager.DetectUsedPackageManagers(args.Path)

	packageManagers := map[string]func(*types.Dependencies) error{
		"npm": getDependencies.GetDependenciesPackageJsonLock,
		"go":  getDependencies.GetDependenciesGoMod,
	}

	for packageManagerFile, packageManager := range usedPackageManagers {
		if args.Verbose {
			fmt.Println("Detected: ", packageManager, " file: ", packageManagerFile)
		}

		dependencies = types.Dependencies{
			HasRiskyFailLicense: false,
			PackageManagerFile:  packageManagerFile,
			CliArguments:        args,
		}

		if handler, exists := packageManagers[packageManager]; exists {
			err := handler(&dependencies)
			if err != nil {
				log.Fatal(err)
			}
			printDependencies.ShowDependencies(&dependencies)
		}

		hasRiskyFailLicense = hasRiskyFailLicense || dependencies.HasRiskyFailLicense

		// Reset dependencies
		dependencies = types.Dependencies{}
	}

	if hasRiskyFailLicense {
		log.Fatal("Risky fail licenses found")
	}
}

/*
Copyright 2023 DigiWolk.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"github.com/digi-wolk/oss-license-auditor/internal/cli"
	detectPackageManager "github.com/digi-wolk/oss-license-auditor/internal/detect/package-manager"
	getDependencies "github.com/digi-wolk/oss-license-auditor/internal/get-dependencies"
	printDependencies "github.com/digi-wolk/oss-license-auditor/internal/print-dependencies"
	"github.com/digi-wolk/oss-license-auditor/internal/types"
	"log"
)

func main() {
	var hasRiskyFailLicense bool
	var dependencies types.Dependencies

	args := cli.GetCliArguments()

	// TODO: In the output, alongside the filename, npm or package-lock.json or both?
	usedPackageManagers := detectPackageManager.DetectUsedPackageManagers(args.Path)

	packageManagers := map[string]func(*types.Dependencies) error{
		"pnpm": getDependencies.GetDependenciesPnpmLock,
		"npm":  getDependencies.GetDependenciesPackageJsonLock,
		"go":   getDependencies.GetDependenciesGoMod,
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

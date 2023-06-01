package lang

import (
	"io/ioutil"
	"log"
)

// DetectUsedPackageManagers Scan the code and detect used package managers
func DetectUsedPackageManagers(pathToCode string) map[string]string {

	packageManagersMap := map[string]string{
		"package-lock.json": "npm",
		"yarn.lock":         "yarn",
		"pom.xml":           "maven",
		"build.gradle":      "gradle",
		"build.sbt":         "sbt",
		"requirements.txt":  "pip",
		"setup.py":          "pip",
		"composer.json":     "composer",
		"go.mod":            "go",
	}

	files, err := ioutil.ReadDir(pathToCode)
	if err != nil {
		log.Fatal(err)
	}

	detectedPackageManagers := make(map[string]string)
	/**
	 * Loop through all files in the directory recursively and add full path of file with package manager
	 * excluding node_modules and .git directories to map if filename matches key in packageManagersMap.
	 */
	for _, file := range files {
		if file.IsDir() {
			// Skip node_modules and .git directories
			if file.Name() == "node_modules" || file.Name() == ".git" {
				continue
			}
			// Recursively call this function for each subdirectory
			for key, value := range DetectUsedPackageManagers(pathToCode + "/" + file.Name()) {
				detectedPackageManagers[key] = value
			}
		} else {
			// Check if filename matches key in packageManagersMap
			if packageManager, ok := packageManagersMap[file.Name()]; ok {
				detectedPackageManagers[pathToCode+"/"+file.Name()] = packageManager
			}
		}
	}

	return detectedPackageManagers
}

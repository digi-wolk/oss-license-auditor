package lang

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

// DetectUsedProgrammingLanguages Scan the code and detect programming languages
func DetectUsedProgrammingLanguages(pathToCode string) []string {
	var detectedLanguages []string

	fileExtensionMap := map[string]string{
		".c":     "C",
		".cpp":   "C++",
		".cs":    "C#",
		".go":    "Go",
		".java":  "Java",
		".js":    "JavaScript",
		".php":   "PHP",
		".py":    "Python",
		".rb":    "Ruby",
		".rs":    "Rust",
		".swift": "Swift",
		".ts":    "TypeScript",
	}

	files, err := ioutil.ReadDir(pathToCode)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fileExtension := filepath.Ext(file.Name())
		detectedLanguage, _ := fileExtensionMap[fileExtension]
		detectedLanguages = append(detectedLanguages, detectedLanguage)
	}
	return detectedLanguages
}

package lang

import "testing"

func arrayContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// Test DetectUsedProgrammingLanguages with one .go file and one .ts file
func TestDetectUsedProgrammingLanguagesWithASingleGoFile(t *testing.T) {
	fixturePath := "../../../test/fixtures/detect-langs-based-on-file-extension"
	detectedLanguages := DetectUsedProgrammingLanguages(fixturePath)
	if len(detectedLanguages) != 2 {
		t.Errorf("DetectUsedProgrammingLanguages was incorrect, got: %d, want: %d.", len(detectedLanguages), 1)
	}
	if arrayContains(detectedLanguages, "Go") == false {
		t.Errorf("DetectUsedProgrammingLanguages was incorrect, got: %s, want: %s.", detectedLanguages[0], "Go")
	}
	// If any of the detectedLanguages array is not equal to TypeScript
	if arrayContains(detectedLanguages, "TypeScript") == false {
		t.Errorf("DetectUsedProgrammingLanguages was incorrect, got: %s, want: %s.", detectedLanguages[1], "TypeScript")
	}
}

package shared

import (
	"os"
	"path/filepath"
	"testing"
)

// TestOpenFileWithRetrySuccess tests that OpenFileWithRetry successfully opens a file that exists
// This is important to ensure the basic functionality works correctly
func TestOpenFileWithRetrySuccess(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "test-file-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	tempFilePath := tempFile.Name()
	defer os.Remove(tempFilePath) // Clean up after the test

	// Close the file so we can reopen it
	tempFile.Close()

	// Get the relative path to the temporary file
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	relPath, err := filepath.Rel(currentDir, tempFilePath)
	if err != nil {
		t.Fatalf("Failed to get relative path: %v", err)
	}

	// Test opening the file with retry
	file, err := OpenFileWithRetry(relPath)
	if err != nil {
		t.Errorf("OpenFileWithRetry failed: %v", err)
	}
	if file == nil {
		t.Error("OpenFileWithRetry returned nil file")
	} else {
		file.Close() // Close the file if it was opened successfully
	}
}

// TestOpenFileWithRetryFailure tests that OpenFileWithRetry fails for a non-existent file
// This is important to ensure the function correctly handles errors
func TestOpenFileWithRetryFailure(t *testing.T) {
	// Test opening a non-existent file
	nonExistentFile := "non-existent-file.txt"
	file, err := OpenFileWithRetry(nonExistentFile)

	// Verify that the function failed
	if err == nil {
		t.Error("OpenFileWithRetry should have failed for a non-existent file")
		if file != nil {
			file.Close()
		}
	}
}

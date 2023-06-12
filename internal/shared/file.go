package shared

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// RetryAttempts represents the number of retry attempts
const RetryAttempts = 3

// RetryDelay represents the delay duration between retries
const RetryDelay = time.Second * 2

func OpenFileWithRetry(filePath string) (*os.File, error) {
	var file *os.File
	var err error

	basePath, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current path: %v", err)
	}
	clean := filepath.Clean(filePath)
	filePath = filepath.Join(basePath, clean)
	for attempt := 1; attempt <= RetryAttempts; attempt++ {
		file, err = os.Open(filePath)
		if err == nil {
			// File opened successfully, break the retry loop
			return file, nil
		}

		log.Printf("Error opening file %s, retrying in %v...", filePath, RetryDelay)
		time.Sleep(RetryDelay)
	}

	return nil, fmt.Errorf("failed to open file %s after %d attempts: %v", filePath, RetryAttempts, err)
}

package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFileInDB(filename string) (*os.File, error) {
	// 1. Ensure "./db/" directory exists (create if it doesn't)
	dbDir := "./db"
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %v", err)
	}

	// 2. Create the file inside "./db/"
	filePath := filepath.Join(dbDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %v", err)
	}

	// Return the file handle (caller should defer file.Close())
	return file, nil
}

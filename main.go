package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// organizeDirectory contains the core logic for scanning and organizing files.
// By separating this from main(), we make our code more modular and testable.
func organizeDirectory(targetDir string) error {
	fmt.Printf("Scanning directory: %s\n", targetDir)
	fmt.Println("--------------------")

	// Read all entries (files and subdirectories) from the target directory.
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		return fmt.Errorf("error reading directory: %w", err)
	}

	fmt.Println("Organizing files...")
	for _, entry := range entries {
		// We only want to process files, so we skip any subdirectories.
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()
		fileExt := filepath.Ext(fileName)

		// Skip files that do not have an extension (e.g., 'Dockerfile').
		if fileExt == "" {
			fmt.Printf(" - Skipping file with no extension: %s\n", fileName)
			continue
		}

		// Create a clean directory name from the extension (e.g., ".txt" -> "txt").
		dirName := strings.ToLower(strings.TrimPrefix(fileExt, "."))
		destDir := filepath.Join(targetDir, dirName)

		// Create the destination directory if it doesn't already exist.
		// os.MkdirAll is safe to call even if the directory is already there.
		if err := os.MkdirAll(destDir, 0755); err != nil {
			log.Printf("Error creating directory %s: %v. Skipping file %s.\n", destDir, err, fileName)
			continue
		}

		// Construct the full original path and the full new path for the file.
		sourcePath := filepath.Join(targetDir, fileName)
		destPath := filepath.Join(destDir, fileName)

		// Move the file from the source path to the destination path.
		if err := os.Rename(sourcePath, destPath); err != nil {
			log.Printf("Error moving file %s to %s: %v\n", sourcePath, destPath, err)
		} else {
			fmt.Printf(" - Moved %s -> %s/\n", fileName, dirName)
		}
	}

	fmt.Println("--------------------")
	fmt.Println("Organization complete.")
	return nil
}

func main() {
	// The main function is now only responsible for handling
	// command-line arguments and calling the core logic function.
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <directory_path>")
		os.Exit(1)
	}
	targetDir := os.Args[1]

	// Call our core logic function and handle any fatal errors.
	if err := organizeDirectory(targetDir); err != nil {
		log.Fatalf("Failed to organize directory: %v", err)
	}
}

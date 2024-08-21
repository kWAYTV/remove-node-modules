package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// removeNodeModules recursively deletes the node_modules directory and its contents
func removeNodeModules(path string) error {
	fmt.Printf("Removing: %s\n", path)
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("failed to remove %s: %w", path, err)
	}
	return nil
}

// findAndRemoveNodeModules traverses the script directory and removes node_modules folders
func findAndRemoveNodeModules(basePath string) error {
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == "node_modules" {
			if err := removeNodeModules(path); err != nil {
				return err
			}
			// Skip the rest of the node_modules folder to avoid walking into it after deletion
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking the path %q: %w", basePath, err)
	}
	return nil
}

func main() {
	// Get the directory where the script is located
	scriptDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current directory: %v\n", err)
		return
	}

	// Find and remove node_modules folders in the script directory
	if err := findAndRemoveNodeModules(scriptDir); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

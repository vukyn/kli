package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Create(filePaths ...string) error {
	if len(filePaths) == 0 {
		return nil
	}

	// Create the file
	errs := make([]error, 0, len(filePaths))
	for _, filePath := range filePaths {
		if err := createFile(filePath, nil); err != nil {
			errs = append(errs, err)
			continue
		}
	}

	return errors.Join(errs...)
}

func Rename(oldPath, newPath string) error {
	if oldPath == "" || newPath == "" {
		return nil
	}

	// Read the old file and write it to the new file
	// then remove the old file
	oldFile, err := os.ReadFile(oldPath)
	if err != nil {
		return err
	}
	if err := createFile(newPath, oldFile); err != nil {
		return err
	}
	if err := removeFile(oldPath); err != nil {
		return err
	}
	return nil
}

func Remove(filePaths ...string) error {
	if len(filePaths) == 0 {
		return nil
	}

	// Remove the file
	errs := make([]error, 0, len(filePaths))
	for _, filePath := range filePaths {
		if err := removeFile(filePath); err != nil {
			errs = append(errs, err)
			continue
		}
	}

	return errors.Join(errs...)
}

func createFile(filePath string, data []byte) error {
	if filePath == "" {
		return nil
	}

	dirPath, _ := filepath.Split(filePath)
	if dirPath != "" { // File inside a directory
		// Create the directory if it does not exist
		if _, err := os.Stat(dirPath); err != nil {
			if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
				return err
			}
		}
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}
	return nil
}

func removeFile(filePath string) error {
	if filePath == "" {
		return nil
	}

	// Check if the file exists
	stat, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	// Check if the file is a directory
	if stat.IsDir() {
		return fmt.Errorf("cannot remove directory: %s", filePath)
	}

	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}

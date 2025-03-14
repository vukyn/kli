package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Create(fileNames ...string) error {
	if len(fileNames) == 0 {
		return nil
	}

	// Create the file
	errs := make([]error, 0, len(fileNames))
	for _, fileName := range fileNames {
		if err := createFile(fileName, nil); err != nil {
			errs = append(errs, err)
			continue
		}
	}

	return errors.Join(errs...)
}

func Rename(oldName, newName string) error {
	if oldName == "" || newName == "" {
		return nil
	}

	// Read the old file and write it to the new file
	// then remove the old file
	oldFile, err := os.ReadFile(oldName)
	if err != nil {
		return err
	}
	if err := createFile(newName, oldFile); err != nil {
		return err
	}
	if err := removeFile(oldName); err != nil {
		return err
	}
	return nil
}

func Remove(fileNames ...string) error {
	if len(fileNames) == 0 {
		return nil
	}

	// Remove the file
	errs := make([]error, 0, len(fileNames))
	for _, fileName := range fileNames {
		if err := removeFile(fileName); err != nil {
			errs = append(errs, err)
			continue
		}
	}

	return errors.Join(errs...)
}

func createFile(fileName string, data []byte) error {
	if fileName == "" {
		return nil
	}

	dir, _ := filepath.Split(fileName)
	if dir != "" { // File inside a directory
		// Create the directory if it does not exist
		if _, err := os.Stat(dir); err != nil {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return err
			}
		}
	}
	if err := os.WriteFile(fileName, data, 0644); err != nil {
		return err
	}
	return nil
}

func removeFile(fileName string) error {
	if fileName == "" {
		return nil
	}

	// Check if the file exists
	stat, err := os.Stat(fileName)
	if err != nil {
		return err
	}

	// Check if the file is a directory
	if stat.IsDir() {
		return fmt.Errorf("cannot remove directory: %s", fileName)
	}

	if err := os.Remove(fileName); err != nil {
		return err
	}
	return nil
}

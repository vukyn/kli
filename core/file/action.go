package file

import (
	"errors"
	"os"
	"path/filepath"
)

func Create(fileNames ...string) error {
	if len(fileNames) == 0 {
		return nil
	}

	// Create the directory if it does not exist
	errs := make([]error, 0, len(fileNames))
	for _, fileName := range fileNames {
		dir, _ := filepath.Split(fileName)
		if dir != "" {
			if _, err := os.Stat(dir); err != nil {
				if err := os.MkdirAll(dir, os.ModePerm); err != nil {
					errs = append(errs, err)
					continue
				}
			}
		}
		if err := os.WriteFile(fileName, []byte{}, 0644); err != nil {
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
	if err := os.WriteFile(newName, oldFile, 0644); err != nil {
		return err
	}
	if err := os.Remove(oldName); err != nil {
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
		if err := os.Remove(fileName); err != nil {
			errs = append(errs, err)
			continue
		}
	}

	return errors.Join(errs...)
}

package file

import (
	"os"
	"path/filepath"
)

func Create(fileName string) error {
	if fileName == "" {
		return nil
	}

	// Create the directory if it does not exist
	dir, _ := filepath.Split(fileName)
	if dir != "" {
		if _, err := os.Stat(dir); err != nil {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return err
			}
		}
	}
	if err := os.WriteFile(fileName, []byte{}, 0644); err != nil {
		return err
	}
	return nil
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

func Remove(fileName string) error {
	if fileName == "" {
		return nil
	}

	// Remove the file
	if err := os.Remove(fileName); err != nil {
		return err
	}
	return nil
}

package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Setup helper function to create a test directory
func setupTestDir(prefix string) string {
	testDir := filepath.Join(os.TempDir(), prefix)
	os.MkdirAll(testDir, os.ModePerm)
	return testDir
}

func TestCreateSingleFile(t *testing.T) {
	testDir := setupTestDir("kli_test_create")
	defer os.RemoveAll(testDir)

	fileName := filepath.Join(testDir, "test1.txt")
	err := Create(fileName)
	assert.NoError(t, err)

	_, err = os.Stat(fileName)
	assert.False(t, os.IsNotExist(err), "File %s was not created", fileName)
}

func TestCreateMultipleFiles(t *testing.T) {
	testDir := setupTestDir("kli_test_create")
	defer os.RemoveAll(testDir)

	fileNames := []string{
		filepath.Join(testDir, "test2.txt"),
		filepath.Join(testDir, "test3.txt"),
	}
	err := Create(fileNames...)
	assert.NoError(t, err)

	for _, fileName := range fileNames {
		_, err := os.Stat(fileName)
		assert.False(t, os.IsNotExist(err), "File %s was not created", fileName)
	}
}

func TestCreateFileInNestedDirectory(t *testing.T) {
	testDir := setupTestDir("kli_test_create")
	defer os.RemoveAll(testDir)

	fileName := filepath.Join(testDir, "nested", "dir", "test4.txt")
	err := Create(fileName)
	assert.NoError(t, err)

	_, err = os.Stat(fileName)
	assert.False(t, os.IsNotExist(err), "File %s was not created", fileName)
}

func TestCreateWithEmptyFileList(t *testing.T) {
	err := Create()
	assert.NoError(t, err)
}

func TestCreateWithEmptyFileName(t *testing.T) {
	err := Create("")
	assert.NoError(t, err)
}

func TestCreateFileWithInvalidPermissions(t *testing.T) {
	// This test only works when run as non-root
	if os.Geteuid() == 0 {
		t.Skip("Skipping test when running as root")
	}

	// Try to create a file in a location that requires root permissions
	fileName := "/root/test_no_permission.txt"
	err := Create(fileName)
	assert.Error(t, err)

	// Clean up if somehow we succeeded
	if err == nil {
		os.Remove(fileName)
	}
}

func TestRenameFileInSameDirectory(t *testing.T) {
	testDir := setupTestDir("kli_test_rename")
	defer os.RemoveAll(testDir)

	testContent := "test content"
	oldName := filepath.Join(testDir, "original.txt")
	newName := filepath.Join(testDir, "renamed.txt")

	// Create the original file
	err := os.WriteFile(oldName, []byte(testContent), 0644)
	assert.NoError(t, err, "Failed to create test file")

	err = Rename(oldName, newName)
	assert.NoError(t, err)

	// Check old file doesn't exist
	_, err = os.Stat(oldName)
	assert.True(t, os.IsNotExist(err), "Old file %s still exists", oldName)

	// Check new file exists
	_, err = os.Stat(newName)
	assert.False(t, os.IsNotExist(err), "New file %s was not created", newName)

	// Check content was preserved
	content, err := os.ReadFile(newName)
	assert.NoError(t, err, "Failed to read new file")
	assert.Equal(t, testContent, string(content), "File content was not preserved")
}

func TestRenameFileToAnotherDirectory(t *testing.T) {
	testDir := setupTestDir("kli_test_rename")
	defer os.RemoveAll(testDir)

	testContent := "test content"
	oldName := filepath.Join(testDir, "source.txt")
	newDir := filepath.Join(testDir, "subdir")
	os.MkdirAll(newDir, os.ModePerm)
	newName := filepath.Join(newDir, "target.txt")

	// Create the original file
	err := os.WriteFile(oldName, []byte(testContent), 0644)
	assert.NoError(t, err, "Failed to create test file")

	err = Rename(oldName, newName)
	assert.NoError(t, err)

	// Check old file doesn't exist
	_, err = os.Stat(oldName)
	assert.True(t, os.IsNotExist(err), "Old file %s still exists", oldName)

	// Check new file exists
	_, err = os.Stat(newName)
	assert.False(t, os.IsNotExist(err), "New file %s was not created", newName)
}

func TestRenameFileWithSpecialCharacters(t *testing.T) {
	testDir := setupTestDir("kli_test_rename")
	defer os.RemoveAll(testDir)

	testContent := "test content"
	oldName := filepath.Join(testDir, "special.txt")
	newName := filepath.Join(testDir, "special!@#$%.txt")

	// Create the original file
	err := os.WriteFile(oldName, []byte(testContent), 0644)
	assert.NoError(t, err, "Failed to create test file")

	err = Rename(oldName, newName)
	assert.NoError(t, err)

	// Check new file exists
	_, err = os.Stat(newName)
	assert.False(t, os.IsNotExist(err), "New file %s was not created", newName)
}

func TestRenameWithEmptyFileNames(t *testing.T) {
	err := Rename("", "")
	assert.NoError(t, err)
}

func TestRenameNonExistentFile(t *testing.T) {
	testDir := setupTestDir("kli_test_rename")
	defer os.RemoveAll(testDir)

	oldName := filepath.Join(testDir, "nonexistent.txt")
	newName := filepath.Join(testDir, "target.txt")

	err := Rename(oldName, newName)
	assert.Error(t, err, "Expected error for non-existent file, got none")
}

func TestRenameToLocationWithoutPermission(t *testing.T) {
	// This test only works when run as non-root
	if os.Geteuid() == 0 {
		t.Skip("Skipping test when running as root")
	}

	testDir := setupTestDir("kli_test_rename")
	defer os.RemoveAll(testDir)

	testContent := "test content"
	oldName := filepath.Join(testDir, "source_perm.txt")
	// Create the original file
	err := os.WriteFile(oldName, []byte(testContent), 0644)
	assert.NoError(t, err, "Failed to create test file")

	// Try to rename to a location that requires root permissions
	newName := "/root/renamed_no_permission.txt"
	err = Rename(oldName, newName)
	assert.Error(t, err, "Expected permission error, got no error")

	// Clean up if somehow we succeeded
	if err == nil {
		os.Remove(newName)
	}
}

func TestRemoveSingleFile(t *testing.T) {
	testDir := setupTestDir("kli_test_remove")
	defer os.RemoveAll(testDir)

	fileName := filepath.Join(testDir, "to_remove.txt")

	// Create the file
	err := os.WriteFile(fileName, []byte("test content"), 0644)
	assert.NoError(t, err, "Failed to create test file")

	err = Remove(fileName)
	assert.NoError(t, err)

	// Check file doesn't exist
	_, err = os.Stat(fileName)
	assert.True(t, os.IsNotExist(err), "File %s still exists", fileName)
}

func TestRemoveMultipleFiles(t *testing.T) {
	testDir := setupTestDir("kli_test_remove")
	defer os.RemoveAll(testDir)

	fileNames := []string{
		filepath.Join(testDir, "multi1.txt"),
		filepath.Join(testDir, "multi2.txt"),
	}

	// Create the files
	for _, fileName := range fileNames {
		err := os.WriteFile(fileName, []byte("test content"), 0644)
		assert.NoError(t, err, "Failed to create test file")
	}

	err := Remove(fileNames...)
	assert.NoError(t, err)

	// Check files don't exist
	for _, fileName := range fileNames {
		_, err := os.Stat(fileName)
		assert.True(t, os.IsNotExist(err), "File %s still exists", fileName)
	}
}

func TestRemoveFileWithSpecialCharacters(t *testing.T) {
	testDir := setupTestDir("kli_test_remove")
	defer os.RemoveAll(testDir)

	fileName := filepath.Join(testDir, "special!@#$%.txt")

	// Create the file
	err := os.WriteFile(fileName, []byte("test content"), 0644)
	assert.NoError(t, err, "Failed to create test file")

	err = Remove(fileName)
	assert.NoError(t, err)

	// Check file doesn't exist
	_, err = os.Stat(fileName)
	assert.True(t, os.IsNotExist(err), "File %s still exists", fileName)
}

func TestRemoveWithEmptyFileList(t *testing.T) {
	err := Remove()
	assert.NoError(t, err)
}

func TestRemoveNonExistentFile(t *testing.T) {
	testDir := setupTestDir("kli_test_remove")
	defer os.RemoveAll(testDir)

	fileName := filepath.Join(testDir, "nonexistent.txt")

	err := Remove(fileName)
	assert.Error(t, err, "Expected error for non-existent file, got none")
}

func TestRemoveDirectory(t *testing.T) {
	testDir := setupTestDir("kli_test_remove")
	defer os.RemoveAll(testDir)

	dirName := filepath.Join(testDir, "test_dir")
	err := os.MkdirAll(dirName, os.ModePerm)
	assert.NoError(t, err)

	err = Remove(dirName)
	assert.Error(t, err, "Expected error when removing directory, got none")
}

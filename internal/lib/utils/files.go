// File: internal/lib/utils/files.go
// Purpose: File utility functions used for checking, reading, and writing files.

package utils

import (
	"errors"
	"io/ioutil"
	"os"
)

// FileExists checks if a given file path exists and is not a directory.
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// DirExists checks if a given path is a directory.
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// ReadFile reads the contents of a file into a string.
func ReadFile(path string) (string, error) {
	if !FileExists(path) {
		return "", errors.New("file does not exist")
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFile writes a string to the specified path, creating it if necessary.
func WriteFile(path, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}

// EnsureDir ensures a directory exists, creating it if it doesn't.
func EnsureDir(path string) error {
	if DirExists(path) {
		return nil
	}
	return os.MkdirAll(path, 0755)
}

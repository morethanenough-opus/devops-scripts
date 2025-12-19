package devops_scripts

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GenerateRandomString generates a random string of a specified length.
func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err!= nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// GetAbsoluteFilePath returns the absolute path of a file.
func GetAbsoluteFilePath(filePath string) (string, error) {
	absPath, err := filepath.Abs(filePath)
	if err!= nil {
		return "", err
	}
	return absPath, nil
}

// IsFileExist checks if a file exists.
func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

// GetFileExtension returns the file extension.
func GetFileExtension(filePath string) string {
	return filepath.Ext(filePath)
}

// GetFileNameWithoutExtension returns the file name without extension.
func GetFileNameWithoutExtension(filePath string) string {
	return strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
}

// GetFileDir returns the directory of a file.
func GetFileDir(filePath string) string {
	return filepath.Dir(filePath)
}

// GetFileSize returns the size of a file in bytes.
func GetFileSize(filePath string) (int64, error) {
 fileInfo, err := os.Stat(filePath)
	if err!= nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

// GetFileModTime returns the modification time of a file.
func GetFileModTime(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err!= nil {
		return 0, err
	}
	return fileInfo.ModTime().Unix(), nil
}

// WriteFile writes data to a file.
func WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

// AppendToFile appends data to a file.
func AppendToFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, append(data, '\n'), 0644)
}

// ReadFile reads the contents of a file.
func ReadFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

// DeleteFile deletes a file.
func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

// CreateDir creates a directory.
func CreateDir(dirPath string) error {
	return os.MkdirAll(dirPath, 0755)
}

// RmDir deletes a directory.
func RmDir(dirPath string) error {
	return os.RemoveAll(dirPath)
}

// CheckIfDirEmpty checks if a directory is empty.
func CheckIfDirEmpty(dirPath string) bool {
	files, err := os.ReadDir(dirPath)
	if err!= nil {
		log.Println(err)
		return true
	}
	return len(files) == 0
}
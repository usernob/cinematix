package filewriter

import (
	"os"
	"path/filepath"
	"strings"
)

const BASE = "storage"

func SplitPath(path string) ([]string, string) {
	listPath := strings.Split(path, "/")
	filename := listPath[len(listPath)-1]

	var filePath []string
	if len(listPath) > 1 {
		filePath = listPath[:len(listPath)-1]
	}
	return filePath, filename
}

func CreateFile(file string, content string) (string, error) {
	subpath, filename := SplitPath(file)
	finalPath := filepath.Join(BASE, filepath.Join(subpath...))

	err := os.MkdirAll(finalPath, 0755)
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(finalPath, filename)
	writeErr := os.WriteFile(filePath, []byte(content), 0644)
	return filePath, writeErr
}

func DeleteFile(file string) (string, error) {
	subpath, filename := SplitPath(file)
	finalpath := filepath.Join(BASE, filepath.Join(subpath...))
	err := os.Remove(filename)
	return finalpath, err
}

func CheckFileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

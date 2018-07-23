package synth

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Scan - returns a slice containing the path and file pointer to each file in dir
func Scan(dir string) ([]*File, error) {
	files := []*File{}

	err := filepath.Walk(dir, func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			files = append(files, &File{path, file})
		}

		return nil
	})

	return files, err
}

// TrimPaths - removes subPath from the Path of every element in files
func TrimPaths(files []*File, subPath string) {
	for _, file := range files {
		file.Path = strings.Replace(file.Path, subPath, "", -1)
	}
}

// CreateFile - creates a file and all folders along the path
func CreateFile(fullPath string) (*os.File, error) {
	// Create any nested folders
	if strings.Count(fullPath, string(os.PathSeparator)) > 1 {
		index := strings.LastIndex(fullPath, string(os.PathSeparator))
		subPath := fullPath[:index]
		os.MkdirAll(subPath, os.ModePerm)
	}

	// Create the file
	filePtr, err := os.Create(fullPath)
	if err != nil {
		log.Printf("error creating file %s\n", fullPath)
		return nil, err
	}

	return filePtr, nil
}

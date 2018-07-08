package synth

import (
	"os"
	"path/filepath"
	"strings"
)

// Scan - returns a slice containing the path to each file in the specified dir
func Scan(dir string) ([]string, error) {
	files := []string{}

	err := filepath.Walk(dir, func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// TrimPaths - returns a copy of paths with subPath removed from every element
func TrimPaths(paths []string, subPath string) []string {
	trimmed := []string{}

	for _, path := range paths {
		t := strings.Replace(path, subPath, "", -1)
		trimmed = append(trimmed, t)
	}

	return trimmed
}

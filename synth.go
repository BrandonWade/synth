package synth

import (
	"os"
	"path/filepath"
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

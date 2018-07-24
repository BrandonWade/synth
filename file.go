package synth

import "os"

// File - the model representing a file on disk
type File struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
	File os.FileInfo
}

// IsEmpty - returns a bool indicating whether a File is empty
func (f *File) IsEmpty() bool {
	return f.Path == "" && f.Size == 0
}

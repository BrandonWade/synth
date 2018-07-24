package synth

import "os"

// File - the model representing a file on disk
type File struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
	File os.FileInfo
}

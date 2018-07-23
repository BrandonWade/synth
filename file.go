package synth

import "os"

// File - the model for a file on disk
type File struct {
	Path string      `json:"path"`
	File os.FileInfo `json:"file"`
}

package openapi

import (
	"io"

	"bazil.org/fuse/fs"
)

func NewFilesystem(openapijson io.Reader) *FileSystem {
	return &FileSystem{}
}

type FileSystem struct{}

func (f *FileSystem) Root() (fs.Node, error) {
	return newDirectory(0, "/")
}

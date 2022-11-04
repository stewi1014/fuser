package openapi

import (
	"context"

	"bazil.org/fuse"
)

func newFile() (*File, error) {
	return &File{}, nil
}

type File struct {
}

func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {

}

func (f *File) ReadAll(ctx context.Context) ([]byte, error) {

}

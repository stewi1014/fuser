package openapi

import (
	"context"
	"os"
	"syscall"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

func newDirectory(parentInode uint64, name string) (*Directory, error) {
	return &Directory{
		inode: fs.GenerateDynamicInode(parentInode, name),
		name:  name,
	}, nil
}

type Directory struct {
	inode uint64
	name  string

	subnodes map[string]fs.Node
	dirent   []fuse.Dirent
}

func (d *Directory) Attr(ctx context.Context, attr *fuse.Attr) error {
	attr.Inode = d.inode
	attr.Mode = os.ModeDir | 0o555
	return nil
}

func (d *Directory) Lookup(ctx context.Context, name string) (fs.Node, error) {
	n, ok := d.subnodes[name]
	if !ok {
		return nil, syscall.ENOENT
	}

	return n, nil
}

func (d *Directory) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	return d.dirent, nil
}

func (d *Directory) add(node fs.Node, name string, ctx context.Context) error {
	attr := fuse.Attr{}
	if err := node.Attr(ctx, &attr); err != nil {
		return err
	}

	var entType fuse.DirentType

	switch node.(type) {
	case *Directory:
		entType = fuse.DT_Dir
	case *File:
		entType = fuse.DT_File
	default:
		entType = fuse.DT_Unknown
	}

	d.subnodes[name] = node
	d.dirent = append(d.dirent, fuse.Dirent{
		Inode: attr.Inode,
		Type:  entType,
		Name:  name,
	})

	return nil
}

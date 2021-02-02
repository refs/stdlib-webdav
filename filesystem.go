package main

import (
	"golang.org/x/net/webdav"
	"os"
	"context"
	"path/filepath"
)

/*
type FileSystem interface {
	Mkdir(ctx context.Context, name string, perm os.FileMode) error
	OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (File, error)
	RemoveAll(ctx context.Context, name string) error
	Rename(ctx context.Context, oldName, newName string) error
	Stat(ctx context.Context, name string) (os.FileInfo, error)
}
 */

type fs struct {
	root string // where the fs is mounted
}

func (f fs) Mkdir(ctx context.Context, name string, perm os.FileMode) error {
	return os.Mkdir(filepath.Join(f.root, name), perm)
}

func (f fs) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (webdav.File, error) {
	panic("implement me")
}

func (f fs) RemoveAll(ctx context.Context, name string) error {
	panic("implement me")
}

func (f fs) Rename(ctx context.Context, oldName, newName string) error {
	panic("implement me")
}

func (f fs) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	return os.Stat(filepath.Join(f.root, name))
}

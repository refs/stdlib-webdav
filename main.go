package main

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"

	"github.com/refs/stdlb-webdav/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/net/webdav"
)

type fs struct {
	root string // fs mount point
}

var defaultMountPath = filepath.Join("vfs")

// newFs initializes a main.fs with the default mount path
// defaultMountPath by default uses where the binary runs.
// this is a test package, so there be dragons.
func newFs() *fs {
	os.MkdirAll(defaultMountPath, 0777)
	return &fs{defaultMountPath}
}

func (f fs) Mkdir(ctx context.Context, name string, perm os.FileMode) error {
	if err := os.Mkdir(filepath.Join(f.root, name), perm); err != nil {
		log.Err(err).Msg("creating directory")
		return err
	}
	return nil
}

func (f fs) OpenFile(ctx context.Context, name string, flag int, perm os.FileMode) (webdav.File, error) {
	file, err := os.OpenFile(filepath.Join(f.root, name), flag, perm)
	if err != nil {
		log.Err(err).Msg("opening file")
		return nil, err
	}
	return file, nil
}

func (f fs) RemoveAll(ctx context.Context, name string) error {
	return os.RemoveAll(name)
}

func (f fs) Rename(ctx context.Context, oldName, newName string) error {
	return os.Rename(oldName, newName)
}

func (f fs) Stat(ctx context.Context, name string) (os.FileInfo, error) {
	return os.Stat(filepath.Join(f.root, name))
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	ls := webdav.NewMemLS()
	h := webdav.Handler{
		FileSystem: *newFs(),
		LockSystem: ls,
	}

	log.Info().Str("service", "webdav").Msg("server listening on localhost:8082")
	if err := http.ListenAndServe(":8082", middleware.Log(&h)); err != nil {
		os.Exit(1)
	}
}

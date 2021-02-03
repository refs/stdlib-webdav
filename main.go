package main

import (
	"context"
	"encoding/xml"
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

// TODO need our own package os.File
// TODO the internal os.File needs to implement webdav.DeadPropsHolder

type file struct {
	os.File
}

func (f *file) DeadProps() (map[xml.Name]webdav.Property, error) {
	var deadProps = make(map[xml.Name]webdav.Property)
	ocNS := xml.Name{
		Space: "http://owncloud.org/ns",
		Local: "B",
	}

	messageP := webdav.Property{
		XMLName: xml.Name{
			Space: "http://owncloud.org/ns",
			Local: "message",
		},
	}
	deadProps[ocNS] = messageP
	return deadProps, nil
}

// TODO continue here. Called in webdav.go:600
func (f *file) Patch(proppatches []webdav.Proppatch) ([]webdav.Propstat, error) {
	panic("implement me")
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

package storage

import (
	"bytes"
)

type File struct {
	Name   string
	Path   string
	Buffer *bytes.Buffer
}

func NewFile(name string, path string) *File {
	return &File{
		Name:   name,
		Path:   path,
		Buffer: &bytes.Buffer{},
	}
}

func (f *File) Write(chunk []byte) error {
	_, err := f.Buffer.Write(chunk)

	return err
}

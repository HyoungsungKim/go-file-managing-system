package storage

import (
	"bytes"
)

type File struct {
	Name   string
	Path   string
	buffer *bytes.Buffer
}

func NewFile(name string, path string) *File {
	return &File{
		Name:   name,
		Path:   path,
		buffer: &bytes.Buffer{},
	}
}

func (f *File) Write(chunk []byte) error {
	_, err := f.buffer.Write(chunk)

	return err
}

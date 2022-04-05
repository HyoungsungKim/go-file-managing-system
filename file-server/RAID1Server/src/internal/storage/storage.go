package storage

import (
	"io/ioutil"
	"log"
	"os"
)

type Manager interface {
	Store(file *File) error
}

type Storage struct {
	dir string
}

func New(dir string) Storage {
	return Storage{
		dir: dir,
	}
}

func (s Storage) Store(file *File) error {
	log.Println("Stored to :", s.dir+file.Path+file.Name)
	if _, err := os.Stat(s.dir + file.Path); err != nil {
		log.Println("Directory does not exist. Generate directory")
		os.Mkdir(s.dir+file.Path, os.ModePerm)
	}

	if err := ioutil.WriteFile(s.dir+file.Path+file.Name, file.buffer.Bytes(), 0644); err != nil {
		log.Println("Write error, ", err)
		return err
	}

	return nil
}

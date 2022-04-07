package storage

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Manager interface {
	Store(file *File) error
	Delete(file *File) error
}

type Storage struct {
	dir string
}

func New(dir string) Storage {
	return Storage{
		dir: dir,
	}
}

func (s Storage) checkExist(file *File) bool {
	if _, err := os.Stat(s.dir + file.Path); err != nil {
		log.Println("Directory does not exist.")
		return false
	} else {
		return true
	}
}

func (s Storage) Store(file *File) error {
	log.Println("Stored to :", s.dir+file.Path+file.Name)
	if !s.checkExist(file) {
		log.Println("Generate directory")
		os.Mkdir(s.dir+file.Path, os.ModePerm)
	}

	if err := ioutil.WriteFile(s.dir+file.Path+file.Name, file.Buffer.Bytes(), 0644); err != nil {
		log.Println("Write error, ", err)
		return err
	}

	return nil
}

func (s Storage) Delete(file *File) error {
	log.Println("Delete :", s.dir+file.Path+file.Name)
	if !s.checkExist(file) {
		return errors.New("THERE IS NO FILE TO DELETE")
	} else {
		os.Remove(s.dir + file.Path + file.Name)
		return nil
	}
}

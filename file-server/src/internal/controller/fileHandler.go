package controller

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type httpStatusCode int
type FileHandler struct {
	*gin.Context
}

// checExist function lets client check file exists from `src`.
// param 'src' denotes source directory of downloading file.
//
// example params: fileName="helloFile.txt", src="storage/users/USER_NAME/"
func (fileHandler *FileHandler) checkExist(fileName string, src string) (httpStatusCode, error) {
	filePath := strings.Join([]string{src, fileName}, "")
	log.Println(filePath)

	if _, err := os.Stat(filePath); err != nil {
		return http.StatusBadRequest, err
	} else {
		return http.StatusOK, nil
	}
}

// create function lets client upload a file to `dst`.
// param `dst` denotes destination directory of created file.
//
// example params: dst="storage/users/USER_NAME/"
func (fileHandler *FileHandler) create(dst string) (httpStatusCode, error) {
	file, _ := fileHandler.FormFile("file")
	filePath := strings.Join([]string{dst, file.Filename}, "")
	log.Println(filePath)

	// If dst directory does not exits, generate
	if _, err := os.Stat(filePath); err == nil {
		return http.StatusBadRequest, errors.New("THIS FILE NAME ALREADY EXISTS")
	} else {
		os.Mkdir(dst, os.ModePerm)
	}

	if err := fileHandler.SaveUploadedFile(file, filePath); err != nil {
		return http.StatusBadRequest, err
	} else {
		return http.StatusOK, nil
	}
}

func update() {

}

func (fileHandler *FileHandler) delete(fileName string, src string) (httpStatusCode, error) {
	filePath := strings.Join([]string{src, fileName}, "")
	log.Println(filePath)

	httpStatusCode, err := fileHandler.checkExist(fileName, src)

	if httpStatusCode == http.StatusOK {
		os.Remove(filePath)
		return httpStatusCode, nil
	} else {
		return httpStatusCode, err
	}
}

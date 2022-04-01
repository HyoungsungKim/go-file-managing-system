package controller

import (
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

// create function lets client upload a file to `dst`.
// param `dst` denotes destination directory of created file.
//
// example params: dst="storage/users/USER_NAME/"
func (fileHandler *FileHandler) create(dst string) (httpStatusCode, error) {
	file, _ := fileHandler.FormFile("file")
	log.Println(file.Filename)

	if _, err := os.Stat(dst); err != nil {
		os.Mkdir(dst, os.ModePerm)
	}

	if err := fileHandler.SaveUploadedFile(file, dst+file.Filename); err != nil {
		return http.StatusBadRequest, err
	} else {
		return http.StatusOK, nil
	}
}

// read function lets client check file exists from `src`.
// param 'src' denotes source directory of downloading file.
//
// example params: fileName="helloFile.txt", src="storage/users/USER_NAME/"
func (fileHander *FileHandler) read(fileName string, src string) (httpStatusCode, error) {
	filePath := strings.Join([]string{src, fileName}, "")
	log.Println(filePath)

	if _, err := os.Stat(filePath); err != nil {
		return http.StatusBadRequest, err
	} else {
		return http.StatusOK, nil
	}
}

func update() {

}

func delete() {

}

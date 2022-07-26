package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"fileServer.com/FileServer/src/internal/gRPCHandler"
	"github.com/gin-gonic/gin"
)

type httpStatusCode int
type FileHandler struct {
	*gin.Context
}

// checkExist function lets client check file exists from `src`.
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
func (fileHandler *FileHandler) create(root string, dirId string, gRPCHandler *gRPCHandler.Client) (httpStatusCode, error) {
	dst := root + "/" + dirId + "/"
	file, _ := fileHandler.FormFile("file")

	filePath := strings.Join([]string{dst, file.Filename}, "")
	log.Println(dst)
	log.Println(filePath)

	// If dst directory does not exits, generate
	if _, err := os.Stat(filePath); err == nil {
		return http.StatusBadRequest, errors.New("THIS FILE NAME ALREADY EXISTS")
	} else {
		log.Println("There is no directory. Generate new direcory")
		log.Println(dst)
		os.Mkdir(dst, os.ModePerm)
	}

	if err := fileHandler.SaveUploadedFile(file, filePath); err != nil {
		return http.StatusBadRequest, err
	} else {

		// Execute gRPC to send a file to RAID1 server
		// Send file info to RAID1 server
		log.Println("Sending file info to RAID1 server...")
		if ack, err := gRPCHandler.SendFileInfo(context.Background(), dirId, file.Filename); err != nil {
			return httpStatusCode(ack.AckStatusCode), errors.New(ack.AckStatusMessage)
		} else {
			log.Println(ack.AckStatusCode, ack.AckStatusMessage)
		}

		// Send file to RAID1 server
		log.Println("Sending file to RAID1 server...")
		if ack, err := gRPCHandler.StreamFile(context.Background(), root, dirId, file.Filename); err != nil {
			return httpStatusCode(ack.AckStatusCode), errors.New(ack.AckStatusMessage)
		} else {
			return http.StatusOK, nil
		}
	}
}

func update() {

}

func (fileHandler *FileHandler) delete(root string, dirId string, fileName string, gRPCHandler *gRPCHandler.Client) (httpStatusCode, error) {
	src := root + "/" + dirId + "/"
	filePath := strings.Join([]string{src, fileName}, "")
	log.Println(filePath)

	statusCode, err := fileHandler.checkExist(fileName, src)

	if statusCode == http.StatusOK {
		if ack, err := gRPCHandler.SendFileInfo(context.Background(), dirId, fileName); err != nil {
			return httpStatusCode(ack.AckStatusCode), errors.New(ack.AckStatusMessage)
		} else {
			log.Println(ack.AckStatusCode, ack.AckStatusMessage)
		}

		if ack, err := gRPCHandler.DeleteFile(context.Background(), root, dirId, fileName); err != nil {
			log.Println(ack.AckStatusCode)
			log.Println(ack.AckStatusMessage)
			return httpStatusCode(ack.AckStatusCode), errors.New(ack.AckStatusMessage)
		} else {
			os.Remove(filePath)
			return statusCode, nil
		}
	}
	return statusCode, err
}

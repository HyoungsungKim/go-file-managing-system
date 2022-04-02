package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var readableType = map[string]bool{
	"txt": true,
	"png": true,
	"gif": true,
	"pdf": true,
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	/* Check server is connected
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	*/

	router.GET("/uploadpage", func(c *gin.Context) {
		title := "upload single file"
		c.HTML(http.StatusOK, "index.html", gin.H{
			"page": title,
		})
	})

	// upload file to UUID. UUID denotes directory
	// For example, storage/1v5bor...0B/`FILE_NAME`
	router.POST("/upload/:UUID", func(c *gin.Context) {
		fileHandler := FileHandler{c}
		dst := "./storage/" + c.Param("UUID") + "/"

		// type of httpStatusCode is int
		httpStatusCode, err := fileHandler.create(dst)

		if err != nil {
			c.String(int(httpStatusCode), fmt.Sprintf("Status code: %d, Upload file err: %s\n", int(httpStatusCode), err.Error()))
		} else {
			c.String(int(httpStatusCode), fmt.Sprintf("Status code: %d, File is uploaded!\n", int(httpStatusCode)))
		}
	})

	// Download file from storage
	router.GET("/download/:UUID/:filename", func(c *gin.Context) {
		fileHandler := FileHandler{c}
		src := "./storage/" + c.Param("UUID") + "/"
		fileName := c.Param("filename")

		httpStatusCode, err := fileHandler.checkExist(fileName, src)

		if err != nil {
			log := fmt.Sprintf("%s does not exists. Please check your file name again. \n", fileName)
			c.String(int(httpStatusCode), log)
		} else {
			c.Writer.WriteHeader(int(httpStatusCode))
			// These headers prevent corrupting extension
			// Also, these headers let a browser download a contents not just show on browser
			c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
			c.Header("Content-Type", "application/octet-stream")
			c.File(src + fileName)
		}
	})

	// Read file from storage and show
	router.GET("/view/:UUID/:filename", func(c *gin.Context) {
		fileHandler := FileHandler{c}
		src := "./storage/" + c.Param("UUID") + "/"
		fileName := c.Param("filename")
		fileExtension := fileName[strings.LastIndex(fileName, ".")+1:]

		httpStatusCode, err := fileHandler.checkExist(fileName, src)

		if err != nil {
			log := fmt.Sprintf("%s does not exists. Please check your file name again. \n", fileName)
			c.String(int(httpStatusCode), log)
		} else {
			if !readableType[fileExtension] {
				log := fmt.Sprintf("%s exists, but browser cannot open up this file extension. Please download this file \n", fileName)
				c.String(int(httpStatusCode), log)
			} else {
				c.Writer.WriteHeader(int(httpStatusCode))
				c.File(src + fileName)
			}
		}
	})

	router.DELETE("/view/:UUID/:filename", func(c *gin.Context) {
		fileHandler := FileHandler{c}
		src := "./storage/" + c.Param("UUID") + "/"
		fileName := c.Param("filename")

		httpStatusCode, err := fileHandler.delete(fileName, src)

		if err != nil {
			log := fmt.Sprintf("%s does not exists. Please check your file name again. \n", fileName)
			c.String(int(httpStatusCode), log)
		} else {
			log := fmt.Sprintf("%s is deleted.", fileName)
			c.String(int(httpStatusCode), log)
		}

	})

	return router
}

package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	router.GET("/download/:UUID/:filename", func(c *gin.Context) {
		fileHandler := FileHandler{c}
		src := "./storage/" + c.Param("UUID") + "/"
		fileName := c.Param("filename")

		// If filename is "" then return bad request code
		log.Println(src)
		log.Println(fileName)
		log.Println(src + fileName)

		httpStatusCode, err := fileHandler.read(fileName, src)

		if err != nil {
			log := fmt.Sprintf("%s does not exists. Please check your file name again. \n", fileName)
			c.String(int(httpStatusCode), log)
		} else {
			//log := fmt.Sprintf("%s exists. \n", fileName)

			/*
				c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
				c.Header("Content-Type", "application/octet-stream")
				c.Header("Content-Length", "1024")
			*/
			c.Writer.WriteHeader(int(httpStatusCode))
			// These headers prevent corrupting extension
			// Also, these headers let a browser download a contents not just show on browser
			c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
			c.Header("Content-Type", "application/octet-stream")
			c.File(src + fileName)
		}

	})

	return router
}

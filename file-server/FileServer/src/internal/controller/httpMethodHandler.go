package controller

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"fileServer.com/FileServer/src/internal/gRPCHandler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var readableType = map[string]bool{
	"txt": true,
	"png": true,
	"gif": true,
	"pdf": true,
	"mp3": true,
	"wmv": true,
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// CORS access
	router.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{
				"http://localhost:3001", // for dev
				"http://localhost:3010",
				"http://172.30.0.1:8090",
				"http://172.30.0.1:3000",
			},
			AllowMethods: []string{"GET", "POST"},
			MaxAge:       12 * time.Hour,
		},
	))

	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	root := "../storage"
	if err != nil {
		log.Fatalln(err)
	}
	//defer conn.Close()

	client := gRPCHandler.NewClient(conn)

	// Check server is connected
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Fileserver pong",
		})
	})

	router.GET("/upload", func(c *gin.Context) {
		title := "upload single file"
		c.HTML(http.StatusOK, "index.html", gin.H{
			"page": title,
		})
	})

	// Upload file to dirID. dirID denotes directory
	// root/accountID/FILE_NAME
	// For example, storage/1v5bor...0B/`FILE_NAME`
	router.POST("/upload/:dirID", func(c *gin.Context) {
		fileHandler := FileHandler{c}
		dirID := c.Param("dirID")

		// type of httpStatusCode is int
		httpStatusCode, err := fileHandler.create(root, dirID, &client)

		if err != nil {
			c.String(int(httpStatusCode), fmt.Sprintf("Status code: %d, Upload file err: %s\n", int(httpStatusCode), err.Error()))
		} else {
			c.String(int(httpStatusCode), fmt.Sprintf("Status code: %d, File is uploaded!\n", int(httpStatusCode)))
		}
	})

	// Upload metadata file to dirID.
	router.POST("/upload/:dirID/metadata", func(c *gin.Context) {

		fileHandler := FileHandler{c}
		dirID := c.Param("dirID")

		// type of httpStatusCode is int
		httpStatusCode, err := fileHandler.create(root, dirID, &client)

		if err != nil {
			c.String(int(httpStatusCode), fmt.Sprintf("Status code: %d, Upload file err: %s\n", int(httpStatusCode), err.Error()))
		} else {
			c.String(int(httpStatusCode), fmt.Sprintf("Status code: %d, File is uploaded!\n", int(httpStatusCode)))
		}
	})

	// Download file from storage
	router.GET("/download/:dirID/:filename", func(c *gin.Context) {
		fileHandler := FileHandler{c}
		src := root + "/" + c.Param("dirID") + "/"
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
	router.GET("/collection/:dirID/:filename", func(c *gin.Context) {
		fileHandler := FileHandler{c}
		src := root + "/" + c.Param("dirID") + "/"
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

	router.DELETE("/collection/:dirID/:filename", func(c *gin.Context) {
		fileHandler := FileHandler{c}
		dirId := c.Param("dirID")
		fileName := c.Param("filename")

		httpStatusCode, err := fileHandler.delete(root, dirId, fileName, &client)

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
